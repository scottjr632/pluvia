package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func previewStack(ctx context.Context, stack *auto.Stack) error {
	outs, err := stack.Outputs(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Stack outputs:")
	for key, value := range outs {
		fmt.Printf("  %v: %v\n", key, value.Value)
	}
	return nil
}

func createInfrastructure(ctx *pulumi.Context) error {
	securityGroup, err := ec2.NewSecurityGroup(ctx, "my-security-group", &ec2.SecurityGroupArgs{
		Description: pulumi.String("Enable SSH access"),
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
	})
	if err != nil {
		return err
	}

	instance, err := ec2.NewInstance(ctx, "my-ec2-instance", &ec2.InstanceArgs{
		InstanceType:        pulumi.String("t2.micro"),
		Ami:                 pulumi.String("ami-0c55b159cbfafe1f0"), // Amazon Linux 2 AMI in us-east-1
		VpcSecurityGroupIds: pulumi.StringArray{securityGroup.ID()},
	})
	if err != nil {
		return err
	}

	// Export the instance's public IP and the security group ID
	ctx.Export("instancePublicIp", instance.PublicIp)
	ctx.Export("instanceId", instance.ID())
	ctx.Export("securityGroupId", securityGroup.ID())
	return nil
}

// func setConfigPassphrase() {
// 	if os.Getenv("PULUMI_CONFIG_PASSPHRASE") == "" {
// 		os.Setenv("PULUMI_CONFIG_PASSPHRASE", "your-secure-passphrase")
// 	}
// }

func main_test() {
	// setConfigPassphrase()

	ctx := context.Background()

	projectName := "my-aws-project"
	stackName := "dev"

	localBackendURL := "file://~/.pulumi"
	project := workspace.Project{
		Name:    tokens.PackageName(projectName),
		Runtime: workspace.NewProjectRuntimeInfo("go", nil),
		Backend: &workspace.ProjectBackend{
			URL: localBackendURL,
		},
	}

	ws, err := auto.NewLocalWorkspace(ctx, auto.Project(project))
	if err != nil {
		fmt.Printf("Failed to create local workspace: %v\n", err)
		os.Exit(1)
	}

	stack, err := auto.UpsertStack(ctx, stackName, ws)
	if err != nil {
		fmt.Printf("Failed to create stack: %v\n", err)
		os.Exit(1)
	}

	stack.Workspace().SetProgram(createInfrastructure)

	// Set the AWS region
	err = stack.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-east-1"})
	if err != nil {
		fmt.Printf("Failed to set AWS region: %v\n", err)
		os.Exit(1)
	}

	// Preview the stack
	res, err := stack.Preview(ctx)
	if err != nil {
		fmt.Printf("Failed to preview stack: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Preview completed! Resources:\n%v\n", res.StdOut)

	// Print the outputs
	outs, err := stack.Outputs(ctx)
	if err != nil {
		fmt.Printf("Failed to get stack outputs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Stack outputs:")
	for key, value := range outs {
		fmt.Printf("  %v: %v\n", key, value.Value)
	}
}
