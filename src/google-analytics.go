package ga

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/OzqurYalcin/google-analytics/config"
	"github.com/google/go-querystring/query"
)

type API struct {
	sync.Mutex
}

type Client struct {
	ProtocolVersion string `url:"v,omitempty"`
	TrackingID      string `url:"tid,omitempty"`
	AnonymizeIP     bool   `url:"aip,omitempty"`
	DataSource      string `url:"ds,omitempty"`
	QueueTime       int64  `url:"qt,omitempty"`
	CacheBuster     string `url:"z,omitempty"`

	ClientID string `url:"cid,omitempty"`
	UserID   string `url:"uid,omitempty"`

	SessionControl       string `url:"sc,omitempty"`
	IPOverride           string `url:"uip,omitempty"`
	UserAgentOverride    string `url:"ua,omitempty"`
	GeographicalOverride string `url:"geoid,omitempty"`

	DocumentReferrer   string `url:"dr,omitempty"`
	CampaignName       string `url:"cn,omitempty"`
	CampaignSource     string `url:"cs,omitempty"`
	CampaignMedium     string `url:"cm,omitempty"`
	CampaignKeyword    string `url:"ck,omitempty"`
	CampaignContent    string `url:"cc,omitempty"`
	CampaignID         string `url:"ci,omitempty"`
	GoogleAdWordsID    string `url:"gclid,omitempty"`
	GoogleDisplayAdsID string `url:"dclid,omitempty"`

	ScreenResolution string `url:"sr,omitempty"`
	ViewportSize     string `url:"vp,omitempty"`
	DocumentEncoding string `url:"de,omitempty"`
	ScreenColors     string `url:"sd,omitempty"`
	UserLanguage     string `url:"ul,omitempty"`
	JavaEnabled      bool   `url:"je,omitempty"`
	FlashVersion     string `url:"fl,omitempty"`

	HitType           string `url:"t,omitempty"`
	NonInteractionHit bool   `url:"ni,omitempty"`

	DocumentLocationURL string   `url:"dl,omitempty"`
	DocumentHostName    string   `url:"dh,omitempty"`
	DocumentPath        string   `url:"dp,omitempty"`
	DocumentTitle       string   `url:"dt,omitempty"`
	ScreenName          string   `url:"cd,omitempty"`
	ContentGroup        []string `url:"cg<gi>,omitempty"`
	LinkID              string   `url:"linkid,omitempty"`

	ApplicationName        string `url:"an,omitempty"`
	ApplicationID          string `url:"aid,omitempty"`
	ApplicationVersion     string `url:"av,omitempty"`
	ApplicationInstallerID string `url:"aiid,omitempty"`

	EventCategory string `url:"ec,omitempty"`
	EventAction   string `url:"ea,omitempty"`
	EventLabel    string `url:"el,omitempty"`
	EventValue    int64  `url:"ev,omitempty"`

	TransactionID          string  `url:"ti,omitempty"`
	TransactionAffiliation string  `url:"ta,omitempty"`
	TransactionRevenue     float64 `url:"tr,omitempty"`
	TransactionShipping    float64 `url:"ts,omitempty"`
	TransactionTax         float64 `url:"tt,omitempty"`
	TransactionCouponCode  float64 `url:"tcc,omitempty"`

	ItemName     string  `url:"in,omitempty"`
	ItemPrice    float64 `url:"ip,omitempty"`
	ItemQuantity int64   `url:"iq,omitempty"`
	ItemCode     string  `url:"ic,omitempty"`
	ItemCategory string  `url:"iv,omitempty"`

	ProductSKU             []string  `url:"pr<pi>id,omitempty"`
	ProductName            []string  `url:"pr<pi>nm,omitempty"`
	ProductBrand           []string  `url:"pr<pi>br,omitempty"`
	ProductCategory        []string  `url:"pr<pi>ca,omitempty"`
	ProductVariant         []string  `url:"pr<pi>va,omitempty"`
	ProductPrice           []float64 `url:"pr<pi>pr,omitempty"`
	ProductQuantity        []int64   `url:"pr<pi>qt,omitempty"`
	ProductCouponCode      []string  `url:"pr<pi>cc,omitempty"`
	ProductPosition        []int64   `url:"pr<pi>ps,omitempty"`
	ProductCustomDimension []string  `url:"pr<pi>cd<di>,omitempty"`
	ProductCustomMetric    []int64   `url:"pr<oi>cm<mi>,omitempty"`
	ProductAction          string    `url:"pa,omitempty"`
	ProductActionList      string    `url:"pal,omitempty"`

	ProductImpressionListName        []string  `url:"il<li>nm,omitempty"`
	ProductImpressionSKU             []string  `url:"il<li>pi<pi>id,omitempty"`
	ProductImpressionName            []string  `url:"il<li>pi<pi>nm,omitempty"`
	ProductImpressionBrand           []string  `url:"il<li>pi<pi>br,omitempty"`
	ProductImpressionCategory        []string  `url:"il<li>pi<pi>ca,omitempty"`
	ProductImpressionVariant         []string  `url:"il<li>pi<pi>va,omitempty"`
	ProductImpressionPosition        []int64   `url:"il<li>pi<pi>ps,omitempty"`
	ProductImpressionPrice           []float64 `url:"il<li>pi<pi>pr,omitempty"`
	ProductImpressionCustomDimension []string  `url:"il<li>pi<pi>cd<di>,omitempty"`
	ProductImpressionCustomMetric    []int64   `url:"il<li>pi<pi>cm<mi>,omitempty"`

	CheckoutStep       int64  `url:"cos,omitempty"`
	CheckoutStepOption string `url:"col,omitempty"`
	CurrencyCode       string `url:"cu,omitempty"`

	PromotionID       []string `url:"promo<pi>id,omitempty"`
	PromotionName     []string `url:"promo<pi>nm,omitempty"`
	PromotionCreative []string `url:"promo<pi>cr,omitempty"`
	PromotionPosition []string `url:"promo<pi>ps,omitempty"`
	PromotionAction   string   `url:"promoa,omitempty"`

	SocialNetwork      string `url:"sn,omitempty"`
	SocialAction       string `url:"sa,omitempty"`
	SocialActionTarget string `url:"st,omitempty"`

	UserTimingCategory     string `url:"utc,omitempty"`
	UserTimingVariableName string `url:"utv,omitempty"`
	UserTimingTime         int64  `url:"utt,omitempty"`
	UserTimingLabel        string `url:"utl,omitempty"`

	PageLoadTime         int64 `url:"plt,omitempty"`
	PageDownloadTime     int64 `url:"pdt,omitempty"`
	DNSTime              int64 `url:"dns,omitempty"`
	RedirectResponseTime int64 `url:"rrt,omitempty"`
	TCPConnectTime       int64 `url:"tcp,omitempty"`
	ServerResponseTime   int64 `url:"srt,omitempty"`
	DOMInteractiveTime   int64 `url:"dit,omitempty"`
	ContentLoadTime      int64 `url:"clt,omitempty"`

	ExceptionDescription string `url:"exd,omitempty"`
	IsExceptionFatal     bool   `url:"exf,omitempty"`

	CustomDimension []string `url:"cd<di>,omitempty"`
	CustomMetric    []int64  `url:"cm<mi>,omitempty"`

	ExperimentID      string `url:"xid,omitempty"`
	ExperimentVariant string `url:"xvar,omitempty"`
}

func (api *API) Send(client *Client) string {
	apidata, _ := query.Values(client)
	res, err := http.Post(config.ApiUrl, "application/x-www-form-urlencoded", strings.NewReader(apidata.Encode()))
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(response)
}
