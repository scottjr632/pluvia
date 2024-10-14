package securitygroup

import (
	"github.com/pluvia/pluvia/context"
	"github.com/pluvia/pluvia/options"
	"github.com/pluvia/pluvia/result"
	"github.com/pluvia/pluvia/templates"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityGroup struct {
	name        string
	description string
	includeSSH  bool
	other       []*ec2.SecurityGroupIngressArgs

	sg *ec2.SecurityGroup

	ingress ec2.SecurityGroupIngressArray

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

func New(
	name string,
	opts ...options.OptionFn[*SecurityGroup]) (res result.Result[*SecurityGroup]) {

	s := &SecurityGroup{name, "", false, []*ec2.SecurityGroupIngressArgs{}, nil, nil, false}
	options.Apply(s, opts...)

	if s.includeSSH {
		s.other = append(s.other, newSSH())
	}

	var ingress ec2.SecurityGroupIngressArray
	s.ingress = ingress
	for _, o := range s.other {
		ingress = append(ingress, o)
	}

	return result.New(s, nil)
}

func (s *SecurityGroup) Create(ctx *context.Context) error {
	sg, err := ec2.NewSecurityGroup(ctx.Pulumi(), s.name, &ec2.SecurityGroupArgs{
		Description: pulumi.String(s.description),
		Ingress:     s.ingress,
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

var _ templates.Template = (*SecurityGroup)(nil)
