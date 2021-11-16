package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OrganizationalBrandingProperties struct {
    Entity
    // Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
    backgroundColor *string;
    // Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
    backgroundImage []byte;
    // A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    backgroundImageRelativeUrl *string;
    // A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
    bannerLogo []byte;
    // A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
    bannerLogoRelativeUrl *string;
    // A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
    cdnList []string;
    // Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
    signInPageText *string;
    // A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
    squareLogo []byte;
    // A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    squareLogoRelativeUrl *string;
    // String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
    usernameHintText *string;
}
// Instantiates a new organizationalBrandingProperties and sets the default values.
func NewOrganizationalBrandingProperties()(*OrganizationalBrandingProperties) {
    m := &OrganizationalBrandingProperties{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) GetBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundColor
    }
}
// Gets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) GetBackgroundImage()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.backgroundImage
    }
}
// Gets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundImageRelativeUrl
    }
}
// Gets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetBannerLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.bannerLogo
    }
}
// Gets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bannerLogoRelativeUrl
    }
}
// Gets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) GetCdnList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.cdnList
    }
}
// Gets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) GetSignInPageText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInPageText
    }
}
// Gets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.squareLogo
    }
}
// Gets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.squareLogoRelativeUrl
    }
}
// Gets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) GetUsernameHintText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usernameHintText
    }
}
// The deserialization information for the current model
func (m *OrganizationalBrandingProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["backgroundColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["backgroundImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImage(val)
        }
        return nil
    }
    res["backgroundImageRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImageRelativeUrl(val)
        }
        return nil
    }
    res["bannerLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogo(val)
        }
        return nil
    }
    res["bannerLogoRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogoRelativeUrl(val)
        }
        return nil
    }
    res["cdnList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCdnList(res)
        }
        return nil
    }
    res["signInPageText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInPageText(val)
        }
        return nil
    }
    res["squareLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogo(val)
        }
        return nil
    }
    res["squareLogoRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoRelativeUrl(val)
        }
        return nil
    }
    res["usernameHintText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameHintText(val)
        }
        return nil
    }
    return res
}
func (m *OrganizationalBrandingProperties) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OrganizationalBrandingProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("backgroundImage", m.GetBackgroundImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("backgroundImageRelativeUrl", m.GetBackgroundImageRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("bannerLogo", m.GetBannerLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bannerLogoRelativeUrl", m.GetBannerLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("cdnList", m.GetCdnList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInPageText", m.GetSignInPageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogo", m.GetSquareLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoRelativeUrl", m.GetSquareLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameHintText", m.GetUsernameHintText())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
// Parameters:
//  - value : Value to set for the backgroundColor property.
func (m *OrganizationalBrandingProperties) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// Sets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
// Parameters:
//  - value : Value to set for the backgroundImage property.
func (m *OrganizationalBrandingProperties) SetBackgroundImage(value []byte)() {
    m.backgroundImage = value
}
// Sets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the backgroundImageRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(value *string)() {
    m.backgroundImageRelativeUrl = value
}
// Sets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// Parameters:
//  - value : Value to set for the bannerLogo property.
func (m *OrganizationalBrandingProperties) SetBannerLogo(value []byte)() {
    m.bannerLogo = value
}
// Sets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the bannerLogoRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(value *string)() {
    m.bannerLogoRelativeUrl = value
}
// Sets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
// Parameters:
//  - value : Value to set for the cdnList property.
func (m *OrganizationalBrandingProperties) SetCdnList(value []string)() {
    m.cdnList = value
}
// Sets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
// Parameters:
//  - value : Value to set for the signInPageText property.
func (m *OrganizationalBrandingProperties) SetSignInPageText(value *string)() {
    m.signInPageText = value
}
// Sets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
// Parameters:
//  - value : Value to set for the squareLogo property.
func (m *OrganizationalBrandingProperties) SetSquareLogo(value []byte)() {
    m.squareLogo = value
}
// Sets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the squareLogoRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(value *string)() {
    m.squareLogoRelativeUrl = value
}
// Sets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
// Parameters:
//  - value : Value to set for the usernameHintText property.
func (m *OrganizationalBrandingProperties) SetUsernameHintText(value *string)() {
    m.usernameHintText = value
}
