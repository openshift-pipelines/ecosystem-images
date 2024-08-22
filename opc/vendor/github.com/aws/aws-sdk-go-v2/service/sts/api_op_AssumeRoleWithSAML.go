// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of temporary security credentials for users who have been
// authenticated via a SAML authentication response. This operation provides a
// mechanism for tying an enterprise identity store or directory to role-based
// Amazon Web Services access without user-specific credentials or configuration.
// For a comparison of AssumeRoleWithSAML with the other API operations that
// produce temporary credentials, see [Requesting Temporary Security Credentials]and [Comparing the Amazon Web Services STS API operations] in the IAM User Guide.
//
// The temporary security credentials returned by this operation consist of an
// access key ID, a secret access key, and a security token. Applications can use
// these temporary security credentials to sign calls to Amazon Web Services
// services.
//
// # Session Duration
//
// By default, the temporary security credentials created by AssumeRoleWithSAML
// last for one hour. However, you can use the optional DurationSeconds parameter
// to specify the duration of your session. Your role session lasts for the
// duration that you specify, or until the time specified in the SAML
// authentication response's SessionNotOnOrAfter value, whichever is shorter. You
// can provide a DurationSeconds value from 900 seconds (15 minutes) up to the
// maximum session duration setting for the role. This setting can have a value
// from 1 hour to 12 hours. To learn how to view the maximum value for your role,
// see [View the Maximum Session Duration Setting for a Role]in the IAM User Guide. The maximum session duration limit applies when you
// use the AssumeRole* API operations or the assume-role* CLI commands. However
// the limit does not apply when you use those operations to create a console URL.
// For more information, see [Using IAM Roles]in the IAM User Guide.
//
// [Role chaining]limits your CLI or Amazon Web Services API role session to a maximum of one
// hour. When you use the AssumeRole API operation to assume a role, you can
// specify the duration of your role session with the DurationSeconds parameter.
// You can specify a parameter value of up to 43200 seconds (12 hours), depending
// on the maximum session duration setting for your role. However, if you assume a
// role using role chaining and provide a DurationSeconds parameter value greater
// than one hour, the operation fails.
//
// # Permissions
//
// The temporary security credentials created by AssumeRoleWithSAML can be used to
// make API calls to any Amazon Web Services service with the following exception:
// you cannot call the STS GetFederationToken or GetSessionToken API operations.
//
// (Optional) You can pass inline or managed [session policies] to this operation. You can pass a
// single JSON policy document to use as an inline session policy. You can also
// specify up to 10 managed policy Amazon Resource Names (ARNs) to use as managed
// session policies. The plaintext that you use for both inline and managed session
// policies can't exceed 2,048 characters. Passing policies to this operation
// returns new temporary credentials. The resulting session's permissions are the
// intersection of the role's identity-based policy and the session policies. You
// can use the role's temporary credentials in subsequent Amazon Web Services API
// calls to access resources in the account that owns the role. You cannot use
// session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see [Session Policies]in the IAM User Guide.
//
// Calling AssumeRoleWithSAML does not require the use of Amazon Web Services
// security credentials. The identity of the caller is validated by using keys in
// the metadata document that is uploaded for the SAML provider entity for your
// identity provider.
//
// Calling AssumeRoleWithSAML can result in an entry in your CloudTrail logs. The
// entry includes the value in the NameID element of the SAML assertion. We
// recommend that you use a NameIDType that is not associated with any personally
// identifiable information (PII). For example, you could instead use the
// persistent identifier ( urn:oasis:names:tc:SAML:2.0:nameid-format:persistent ).
//
// # Tags
//
// (Optional) You can configure your IdP to pass attributes into your SAML
// assertion as session tags. Each session tag consists of a key name and an
// associated value. For more information about session tags, see [Passing Session Tags in STS]in the IAM User
// Guide.
//
// You can pass up to 50 session tags. The plaintext session tag keys can’t exceed
// 128 characters and the values can’t exceed 256 characters. For these and
// additional limits, see [IAM and STS Character Limits]in the IAM User Guide.
//
// An Amazon Web Services conversion compresses the passed inline session policy,
// managed policy ARNs, and session tags into a packed binary format that has a
// separate limit. Your request can fail for this limit even if your plaintext
// meets the other requirements. The PackedPolicySize response element indicates
// by percentage how close the policies and tags for your request are to the upper
// size limit.
//
// You can pass a session tag with the same key as a tag that is attached to the
// role. When you do, session tags override the role's tags with the same key.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// You can set the session tags as transitive. Transitive tags persist during role
// chaining. For more information, see [Chaining Roles with Session Tags]in the IAM User Guide.
//
// # SAML Configuration
//
// Before your application can call AssumeRoleWithSAML , you must configure your
// SAML identity provider (IdP) to issue the claims required by Amazon Web
// Services. Additionally, you must use Identity and Access Management (IAM) to
// create a SAML provider entity in your Amazon Web Services account that
// represents your identity provider. You must also create an IAM role that
// specifies this SAML provider in its trust policy.
//
// For more information, see the following resources:
//
// [About SAML 2.0-based Federation]
//   - in the IAM User Guide.
//
// [Creating SAML Identity Providers]
//   - in the IAM User Guide.
//
// [Configuring a Relying Party and Claims]
//   - in the IAM User Guide.
//
// [Creating a Role for SAML 2.0 Federation]
//   - in the IAM User Guide.
//
// [View the Maximum Session Duration Setting for a Role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session
// [Creating a Role for SAML 2.0 Federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_saml.html
// [IAM and STS Character Limits]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length
// [Comparing the Amazon Web Services STS API operations]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison
// [Creating SAML Identity Providers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml.html
// [session policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
// [Configuring a Relying Party and Claims]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml_relying-party.html
// [Role chaining]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-role-chaining
// [Using IAM Roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
// [About SAML 2.0-based Federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Chaining Roles with Session Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining
func (c *Client) AssumeRoleWithSAML(ctx context.Context, params *AssumeRoleWithSAMLInput, optFns ...func(*Options)) (*AssumeRoleWithSAMLOutput, error) {
	if params == nil {
		params = &AssumeRoleWithSAMLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssumeRoleWithSAML", params, optFns, c.addOperationAssumeRoleWithSAMLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssumeRoleWithSAMLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssumeRoleWithSAMLInput struct {

	// The Amazon Resource Name (ARN) of the SAML provider in IAM that describes the
	// IdP.
	//
	// This member is required.
	PrincipalArn *string

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// This member is required.
	RoleArn *string

	// The base64 encoded SAML authentication response provided by the IdP.
	//
	// For more information, see [Configuring a Relying Party and Adding Claims] in the IAM User Guide.
	//
	// [Configuring a Relying Party and Adding Claims]: https://docs.aws.amazon.com/IAM/latest/UserGuide/create-role-saml-IdP-tasks.html
	//
	// This member is required.
	SAMLAssertion *string

	// The duration, in seconds, of the role session. Your role session lasts for the
	// duration that you specify for the DurationSeconds parameter, or until the time
	// specified in the SAML authentication response's SessionNotOnOrAfter value,
	// whichever is shorter. You can provide a DurationSeconds value from 900 seconds
	// (15 minutes) up to the maximum session duration setting for the role. This
	// setting can have a value from 1 hour to 12 hours. If you specify a value higher
	// than this setting, the operation fails. For example, if you specify a session
	// duration of 12 hours, but your administrator set the maximum session duration to
	// 6 hours, your operation fails. To learn how to view the maximum value for your
	// role, see [View the Maximum Session Duration Setting for a Role]in the IAM User Guide.
	//
	// By default, the value is set to 3600 seconds.
	//
	// The DurationSeconds parameter is separate from the duration of a console
	// session that you might request using the returned credentials. The request to
	// the federation endpoint for a console sign-in token takes a SessionDuration
	// parameter that specifies the maximum length of the console session. For more
	// information, see [Creating a URL that Enables Federated Users to Access the Amazon Web Services Management Console]in the IAM User Guide.
	//
	// [View the Maximum Session Duration Setting for a Role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session
	// [Creating a URL that Enables Federated Users to Access the Amazon Web Services Management Console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html
	DurationSeconds *int32

	// An IAM policy in JSON format that you want to use as an inline session policy.
	//
	// This parameter is optional. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use the
	// role's temporary credentials in subsequent Amazon Web Services API calls to
	// access resources in the account that owns the role. You cannot use session
	// policies to grant more permissions than those allowed by the identity-based
	// policy of the role that is being assumed. For more information, see [Session Policies]in the IAM
	// User Guide.
	//
	// The plaintext that you use for both inline and managed session policies can't
	// exceed 2,048 characters. The JSON policy characters can be any ASCII character
	// from the space character to the end of the valid character list (\u0020 through
	// \u00FF). It can also include the tab (\u0009), linefeed (\u000A), and carriage
	// return (\u000D) characters.
	//
	// An Amazon Web Services conversion compresses the passed inline session policy,
	// managed policy ARNs, and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates
	// by percentage how close the policies and tags for your request are to the upper
	// size limit.
	//
	// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	Policy *string

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want to
	// use as managed session policies. The policies must exist in the same account as
	// the role.
	//
	// This parameter is optional. You can provide up to 10 managed policy ARNs.
	// However, the plaintext that you use for both inline and managed session policies
	// can't exceed 2,048 characters. For more information about ARNs, see [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]in the
	// Amazon Web Services General Reference.
	//
	// An Amazon Web Services conversion compresses the passed inline session policy,
	// managed policy ARNs, and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates
	// by percentage how close the policies and tags for your request are to the upper
	// size limit.
	//
	// Passing policies to this operation returns new temporary credentials. The
	// resulting session's permissions are the intersection of the role's
	// identity-based policy and the session policies. You can use the role's temporary
	// credentials in subsequent Amazon Web Services API calls to access resources in
	// the account that owns the role. You cannot use session policies to grant more
	// permissions than those allowed by the identity-based policy of the role that is
	// being assumed. For more information, see [Session Policies]in the IAM User Guide.
	//
	// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	// [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	PolicyArns []types.PolicyDescriptorType

	noSmithyDocumentSerde
}

// Contains the response to a successful AssumeRoleWithSAML request, including temporary Amazon Web
// Services credentials that can be used to make Amazon Web Services requests.
type AssumeRoleWithSAMLOutput struct {

	// The identifiers for the temporary security credentials that the operation
	// returns.
	AssumedRoleUser *types.AssumedRoleUser

	//  The value of the Recipient attribute of the SubjectConfirmationData element of
	// the SAML assertion.
	Audience *string

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// The size of the security token that STS API operations return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size.
	Credentials *types.Credentials

	// The value of the Issuer element of the SAML assertion.
	Issuer *string

	// A hash value based on the concatenation of the following:
	//
	//   - The Issuer response value.
	//
	//   - The Amazon Web Services account ID.
	//
	//   - The friendly name (the last part of the ARN) of the SAML provider in IAM.
	//
	// The combination of NameQualifier and Subject can be used to uniquely identify a
	// user.
	//
	// The following pseudocode shows how the hash value is calculated:
	//
	//     BASE64 ( SHA1 ( "https://example.com/saml" + "123456789012" + "/MySAMLIdP" ) )
	NameQualifier *string

	// A percentage value that indicates the packed size of the session policies and
	// session tags combined passed in the request. The request fails if the packed
	// size is greater than 100 percent, which means the policies and tags exceeded the
	// allowed space.
	PackedPolicySize *int32

	// The value in the SourceIdentity attribute in the SAML assertion.
	//
	// You can require users to set a source identity value when they assume a role.
	// You do this by using the sts:SourceIdentity condition key in a role trust
	// policy. That way, actions that are taken with the role are associated with that
	// user. After the source identity is set, the value cannot be changed. It is
	// present in the request for all actions that are taken by the role and persists
	// across [chained role]sessions. You can configure your SAML identity provider to use an
	// attribute associated with your users, like user name or email, as the source
	// identity when calling AssumeRoleWithSAML . You do this by adding an attribute to
	// the SAML assertion. For more information about using source identity, see [Monitor and control actions taken with assumed roles]in
	// the IAM User Guide.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can also
	// include underscores or any of the following characters: =,.@-
	//
	// [chained role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts#iam-term-role-chaining
	// [Monitor and control actions taken with assumed roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html
	SourceIdentity *string

	// The value of the NameID element in the Subject element of the SAML assertion.
	Subject *string

	//  The format of the name ID, as defined by the Format attribute in the NameID
	// element of the SAML assertion. Typical examples of the format are transient or
	// persistent .
	//
	// If the format includes the prefix urn:oasis:names:tc:SAML:2.0:nameid-format ,
	// that prefix is removed. For example,
	// urn:oasis:names:tc:SAML:2.0:nameid-format:transient is returned as transient .
	// If the format includes any other prefix, the format is returned with no
	// modifications.
	SubjectType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssumeRoleWithSAMLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAssumeRoleWithSAML{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAssumeRoleWithSAML{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssumeRoleWithSAML"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpAssumeRoleWithSAMLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssumeRoleWithSAML(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssumeRoleWithSAML(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssumeRoleWithSAML",
	}
}
