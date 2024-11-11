package box

import (
	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/options"
	"github.com/pluvia/pluvia/result"
	"github.com/pluvia/pluvia/templates"
	cloudconfigs "github.com/pluvia/pluvia/templates/cloud-configs"
	"github.com/pluvia/pluvia/templates/securitygroup"
	"github.com/pluvia/pluvia/templates/strategies"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Box struct {
	name          string
	instanceType  string
	ami           string
	includeSSH    bool
	includeDocker bool

	sg       *securitygroup.SecurityGroup
	instance *ec2.Instance

	attached []strategies.Strategy
}

func WithDocker() options.OptionFn[*Box] {
	return func(b *Box) *Box {
		b.includeSSH = true
		b.includeDocker = true
		return b
	}
}

func WithIncludeSSH() options.OptionFn[*Box] {
	return func(b *Box) *Box {
		b.includeSSH = true
		return b
	}
}

func WithSecurityGroup(sg *securitygroup.SecurityGroup) options.OptionFn[*Box] {
	return func(b *Box) *Box {
		b.sg = sg
		return b
	}
}

func New(name string, ami string, instanceType string, opts ...options.OptionFn[*Box]) (res result.Result[*Box]) {
	defer result.Recover(&res)

	b := &Box{name, instanceType, ami, false, false, nil, nil, []strategies.Strategy{}}
	options.Apply(b, opts...)

	if b.includeSSH && b.sg == nil {
		sg := securitygroup.New(name+"-security-group", securitygroup.WithSSH()).Must()
		b.sg = sg
	} else if b.includeSSH && b.sg != nil {
		panic("Cannot include SSH and specify a security group")
	}

	return result.New(b, nil)
}

func (b *Box) Create(ctx *templates.ContextWithPulumi) error {
	if b.sg != nil {
		err := b.sg.Create(ctx)
		if err != nil {
			return err
		}
	}

	ctx.Log().Debug("Creating box " + b.name)
	cConfig := cloudconfigs.NewCloudConfigBuilder(cloudconfigs.WithDockerCloudConfig())

	i, err := ec2.NewInstance(ctx.PL, b.name, &ec2.InstanceArgs{
		InstanceType:        pulumi.String(b.instanceType),
		Ami:                 pulumi.String(b.ami),
		VpcSecurityGroupIds: pulumi.StringArray{b.sg.ID()},
		UserData:            pulumi.String(cConfig.Build()),
	})

	b.instance = i
	return err
}

func (b *Box) Attach(fn strategies.StrategyFn[*Box]) error {
	b.attached = append(b.attached, fn(b))
	return nil
}

func (b *Box) Run(ctx context.Context) error {
	for _, a := range b.attached {
		if a == nil {
			ctx.Log().Error("Found nil strategy attaching to box" + b.name + ", please check strategy attachments.")
			continue
		}

		if err := a.Run(ctx); err != nil {
			return err
		}
	}

	return nil
}

var _ templates.Attachable[*Box] = (*Box)(nil)
