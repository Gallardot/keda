// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Decreases the Kinesis data stream's retention period, which is the length of
// time data records are accessible after they are added to the stream. The minimum
// value of a stream's retention period is 24 hours.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// This operation may result in lost data. For example, if the stream's retention
// period is 48 hours and is decreased to 24 hours, any data already in the stream
// that is older than 24 hours is inaccessible.
func (c *Client) DecreaseStreamRetentionPeriod(ctx context.Context, params *DecreaseStreamRetentionPeriodInput, optFns ...func(*Options)) (*DecreaseStreamRetentionPeriodOutput, error) {
	if params == nil {
		params = &DecreaseStreamRetentionPeriodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecreaseStreamRetentionPeriod", params, optFns, c.addOperationDecreaseStreamRetentionPeriodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecreaseStreamRetentionPeriodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DecreaseStreamRetentionPeriod.
type DecreaseStreamRetentionPeriodInput struct {

	// The new retention period of the stream, in hours. Must be less than the current
	// retention period.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to modify.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *DecreaseStreamRetentionPeriodInput) bindEndpointParams(p *EndpointParameters) {

	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type DecreaseStreamRetentionPeriodOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDecreaseStreamRetentionPeriodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDecreaseStreamRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecreaseStreamRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DecreaseStreamRetentionPeriod"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDecreaseStreamRetentionPeriodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseStreamRetentionPeriod(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDecreaseStreamRetentionPeriod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DecreaseStreamRetentionPeriod",
	}
}
