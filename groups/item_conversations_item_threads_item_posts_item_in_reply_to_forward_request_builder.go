package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder provides operations to call the forward method.
type ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal instantiates a new ForwardRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo/forward", pathParameters),
    }
    return m
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder instantiates a new ForwardRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action forward
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) Post(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action forward
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
