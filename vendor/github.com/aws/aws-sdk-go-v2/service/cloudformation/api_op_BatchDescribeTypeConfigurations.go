// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns configuration data for the specified CloudFormation extensions, from
// the CloudFormation registry for the account and Region. For more information,
// see Configuring extensions at the account level (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration)
// in the CloudFormation User Guide.
func (c *Client) BatchDescribeTypeConfigurations(ctx context.Context, params *BatchDescribeTypeConfigurationsInput, optFns ...func(*Options)) (*BatchDescribeTypeConfigurationsOutput, error) {
	if params == nil {
		params = &BatchDescribeTypeConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDescribeTypeConfigurations", params, optFns, c.addOperationBatchDescribeTypeConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDescribeTypeConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDescribeTypeConfigurationsInput struct {

	// The list of identifiers for the desired extension configurations.
	//
	// This member is required.
	TypeConfigurationIdentifiers []types.TypeConfigurationIdentifier

	noSmithyDocumentSerde
}

type BatchDescribeTypeConfigurationsOutput struct {

	// A list of information concerning any errors generated during the setting of the
	// specified configurations.
	Errors []types.BatchDescribeTypeConfigurationsError

	// A list of any of the specified extension configurations from the CloudFormation
	// registry.
	TypeConfigurations []types.TypeConfigurationDetails

	// A list of any of the specified extension configurations that CloudFormation
	// could not process for any reason.
	UnprocessedTypeConfigurations []types.TypeConfigurationIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDescribeTypeConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchDescribeTypeConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchDescribeTypeConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchDescribeTypeConfigurations"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchDescribeTypeConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDescribeTypeConfigurations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opBatchDescribeTypeConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchDescribeTypeConfigurations",
	}
}
