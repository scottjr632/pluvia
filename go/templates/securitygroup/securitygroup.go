package securitygroup

import (
	"fmt"

	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/options"
	"github.com/pluvia/pluvia/result"
	"github.com/pluvia/pluvia/templates"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityGroup struct {
	name          string
	description   string
	includeSSH    bool
	blockOutbound bool
	other         []*ec2.SecurityGroupIngressArgs

	sg *ec2.SecurityGroup

	ingress ec2.SecurityGroupIngressArray
	egress  ec2.SecurityGroupEgressArray

	isCreated bool
}

type SecurityGroupOption func(*SecurityGroup) *SecurityGroup

func WithSSH() options.OptionFn[*SecurityGroup] {
	return func(s *SecurityGroup) *SecurityGroup {
		s.includeSSH = true
		return s
	}
}

func WithDescription(description string) options.OptionFn[*SecurityGroup] {
	return func(s *SecurityGroup) *SecurityGroup {
		s.description = description
		return s
	}
}

func WithOther(other ...*ec2.SecurityGroupIngressArgs) options.OptionFn[*SecurityGroup] {
	return func(s *SecurityGroup) *SecurityGroup {
		s.other = append(s.other, other...)
		return s
	}
}

func WithBlockOutbound() options.OptionFn[*SecurityGroup] {
	return func(s *SecurityGroup) *SecurityGroup {
		s.blockOutbound = true
		return s
	}
}

func New(
	ctx context.Context,
	name string,
	description string,
	opts ...options.OptionFn[*SecurityGroup]) (res result.Result[*SecurityGroup]) {

	s := &SecurityGroup{
		name:          name,
		description:   description,
		includeSSH:    false,
		blockOutbound: false,
		other:         []*ec2.SecurityGroupIngressArgs{},
		sg:            nil,
		ingress:       nil,
		egress:        nil,
		isCreated:     false,
	}
	options.Apply(s, opts...)

	if s.includeSSH {
		ctx.Log().Debug("Including SSH security group")
		s.ingress = append(s.ingress, newSSH())
	}

	if !s.blockOutbound {
		ctx.Log().Debug("Not blocking outbound traffic")
		s.egress = append(s.egress, newOutboundRule())
	}

	for _, o := range s.other {
		s.ingress = append(s.ingress, o)
	}

	return result.New(s, nil)
}

func (s *SecurityGroup) Create(ctx *templates.ContextWithPulumi) error {
	sg, err := ec2.NewSecurityGroup(ctx.PL, s.name, &ec2.SecurityGroupArgs{
		Description: pulumi.String(s.description),
		Ingress:     s.ingress,
		Egress:      s.egress,
	})
	s.sg = sg

	s.isCreated = true

	return err
}

func (sg *SecurityGroup) ID() pulumi.IDOutput {
	if !sg.isCreated || sg.sg == nil {
		panic("SecurityGroup is not created")
	}
	return sg.sg.ID()
}

func newSSH() *ec2.SecurityGroupIngressArgs {
	return &ec2.SecurityGroupIngressArgs{
		Protocol:   pulumi.String("tcp"),
		FromPort:   pulumi.Int(22),
		ToPort:     pulumi.Int(22),
		CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
	}
}

func newOutboundRule() *ec2.SecurityGroupEgressArgs {
	fmt.Println("newOutboundRule")
	return &ec2.SecurityGroupEgressArgs{
		Protocol:   pulumi.String("-1"),
		FromPort:   pulumi.Int(0),
		ToPort:     pulumi.Int(0),
		CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
	}
}

var _ templates.Template = (*SecurityGroup)(nil)
