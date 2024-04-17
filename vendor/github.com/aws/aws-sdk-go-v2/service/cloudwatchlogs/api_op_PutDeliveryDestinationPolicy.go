// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and assigns an IAM policy that grants permissions to CloudWatch Logs to
// deliver logs cross-account to a specified destination in this account. To
// configure the delivery of logs from an Amazon Web Services service in another
// account to a logs delivery destination in the current account, you must do the
// following:
//   - Create a delivery source, which is a logical object that represents the
//     resource that is actually sending the logs. For more information, see
//     PutDeliverySource (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html)
//     .
//   - Create a delivery destination, which is a logical object that represents
//     the actual delivery destination. For more information, see
//     PutDeliveryDestination (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html)
//     .
//   - Use this operation in the destination account to assign an IAM policy to
//     the destination. This policy allows delivery to that destination.
//   - Create a delivery by pairing exactly one delivery source and one delivery
//     destination. For more information, see CreateDelivery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html)
//     .
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at
// Enabling logging from Amazon Web Services services. (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html)
// The contents of the policy must include two statements. One statement enables
// general logs delivery, and the other allows delivery to the chosen destination.
// See the examples for the needed policies.
func (c *Client) PutDeliveryDestinationPolicy(ctx context.Context, params *PutDeliveryDestinationPolicyInput, optFns ...func(*Options)) (*PutDeliveryDestinationPolicyOutput, error) {
	if params == nil {
		params = &PutDeliveryDestinationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDeliveryDestinationPolicy", params, optFns, c.addOperationPutDeliveryDestinationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDeliveryDestinationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDeliveryDestinationPolicyInput struct {

	// The name of the delivery destination to assign this policy to.
	//
	// This member is required.
	DeliveryDestinationName *string

	// The contents of the policy.
	//
	// This member is required.
	DeliveryDestinationPolicy *string

	noSmithyDocumentSerde
}

type PutDeliveryDestinationPolicyOutput struct {

	// The contents of the policy that you just created.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDeliveryDestinationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutDeliveryDestinationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDeliveryDestinationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutDeliveryDestinationPolicy"); err != nil {
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
	if err = addOpPutDeliveryDestinationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDeliveryDestinationPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDeliveryDestinationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutDeliveryDestinationPolicy",
	}
}
