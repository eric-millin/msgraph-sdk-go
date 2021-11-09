package caseexportoperation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i6c205570953da1ba8a1f0d1b861c2d3158c47ad2c8a1fe8a0fb718daa8cd4b7f "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/caseexportoperation/getdownloadurl"
)

// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.ediscovery.caseExportOperation
type CaseExportOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Instantiates a new CaseExportOperationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    m := &CaseExportOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.ediscovery.caseExportOperation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CaseExportOperationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseExportOperationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.ediscovery.caseExportOperation\microsoft.graph.ediscovery.getDownloadUrl()
func (m *CaseExportOperationRequestBuilder) GetDownloadUrl()(*i6c205570953da1ba8a1f0d1b861c2d3158c47ad2c8a1fe8a0fb718daa8cd4b7f.GetDownloadUrlRequestBuilder) {
    return i6c205570953da1ba8a1f0d1b861c2d3158c47ad2c8a1fe8a0fb718daa8cd4b7f.NewGetDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
