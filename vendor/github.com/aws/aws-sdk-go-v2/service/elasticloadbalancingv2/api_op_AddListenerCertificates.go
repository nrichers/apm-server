// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/AddListenerCertificatesInput
type AddListenerCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The certificate to add. You can specify one certificate per call. Set CertificateArn
	// to the certificate ARN but do not set IsDefault.
	//
	// Certificates is a required field
	Certificates []Certificate `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the listener.
	//
	// ListenerArn is a required field
	ListenerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddListenerCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddListenerCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddListenerCertificatesInput"}

	if s.Certificates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificates"))
	}

	if s.ListenerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/AddListenerCertificatesOutput
type AddListenerCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the certificates in the certificate list.
	Certificates []Certificate `type:"list"`
}

// String returns the string representation
func (s AddListenerCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddListenerCertificates = "AddListenerCertificates"

// AddListenerCertificatesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Adds the specified SSL server certificate to the certificate list for the
// specified HTTPS listener.
//
// If the certificate in already in the certificate list, the call is successful
// but the certificate is not added again.
//
// To get the certificate list for a listener, use DescribeListenerCertificates.
// To remove certificates from the certificate list for a listener, use RemoveListenerCertificates.
// To replace the default certificate for a listener, use ModifyListener.
//
// For more information, see SSL Certificates (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#https-listener-certificates)
// in the Application Load Balancers Guide.
//
//    // Example sending a request using AddListenerCertificatesRequest.
//    req := client.AddListenerCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/AddListenerCertificates
func (c *Client) AddListenerCertificatesRequest(input *AddListenerCertificatesInput) AddListenerCertificatesRequest {
	op := &aws.Operation{
		Name:       opAddListenerCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddListenerCertificatesInput{}
	}

	req := c.newRequest(op, input, &AddListenerCertificatesOutput{})
	return AddListenerCertificatesRequest{Request: req, Input: input, Copy: c.AddListenerCertificatesRequest}
}

// AddListenerCertificatesRequest is the request type for the
// AddListenerCertificates API operation.
type AddListenerCertificatesRequest struct {
	*aws.Request
	Input *AddListenerCertificatesInput
	Copy  func(*AddListenerCertificatesInput) AddListenerCertificatesRequest
}

// Send marshals and sends the AddListenerCertificates API request.
func (r AddListenerCertificatesRequest) Send(ctx context.Context) (*AddListenerCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddListenerCertificatesResponse{
		AddListenerCertificatesOutput: r.Request.Data.(*AddListenerCertificatesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddListenerCertificatesResponse is the response type for the
// AddListenerCertificates API operation.
type AddListenerCertificatesResponse struct {
	*AddListenerCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddListenerCertificates request.
func (r *AddListenerCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}