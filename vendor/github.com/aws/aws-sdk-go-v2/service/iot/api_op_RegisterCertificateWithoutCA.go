// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Register a certificate that does not have a certificate authority (CA). For
// supported certificates, consult  Certificate signing algorithms supported by IoT
// (https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html#x509-cert-algorithms).
func (c *Client) RegisterCertificateWithoutCA(ctx context.Context, params *RegisterCertificateWithoutCAInput, optFns ...func(*Options)) (*RegisterCertificateWithoutCAOutput, error) {
	if params == nil {
		params = &RegisterCertificateWithoutCAInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCertificateWithoutCA", params, optFns, c.addOperationRegisterCertificateWithoutCAMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCertificateWithoutCAOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterCertificateWithoutCAInput struct {

	// The certificate data, in PEM format.
	//
	// This member is required.
	CertificatePem *string

	// The status of the register certificate request.
	Status types.CertificateStatus

	noSmithyDocumentSerde
}

type RegisterCertificateWithoutCAOutput struct {

	// The Amazon Resource Name (ARN) of the registered certificate.
	CertificateArn *string

	// The ID of the registered certificate. (The last part of the certificate ARN
	// contains the certificate ID.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterCertificateWithoutCAMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCertificateWithoutCA{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCertificateWithoutCA{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRegisterCertificateWithoutCAValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCertificateWithoutCA(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opRegisterCertificateWithoutCA(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RegisterCertificateWithoutCA",
	}
}
