package ga

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/OzqurYalcin/google-analytics/config"
)

type API struct {
	sync.Mutex
}

type Product struct {
	SKU             string   `json:"id,omitempty"`
	Name            string   `json:"nm,omitempty"`
	Brand           string   `json:"br,omitempty"`
	Category        string   `json:"ca,omitempty"`
	Variant         string   `json:"va,omitempty"`
	Price           string   `json:"pr,omitempty"`
	Quantity        string   `json:"qt,omitempty"`
	CouponCode      string   `json:"cc,omitempty"`
	Position        string   `json:"ps,omitempty"`
	CustomDimension []string `json:"cd,omitempty"`
	CustomMetric    []string `json:"cm,omitempty"`
	Action          string   `json:"pa,omitempty"`
	ActionList      string   `json:"pal,omitempty"`
}

type ProductImpression struct {
	ListName []string `json:"nm,omitempty"`
	Product  []struct {
		SKU             string   `json:"id,omitempty"`
		Name            string   `json:"nm,omitempty"`
		Brand           string   `json:"br,omitempty"`
		Category        string   `json:"ca,omitempty"`
		Variant         string   `json:"va,omitempty"`
		Position        string   `json:"ps,omitempty"`
		Price           []string `json:"pr,omitempty"`
		CustomDimension []string `json:"cd,omitempty"`
		CustomMetric    []string `json:"cm,omitempty"`
	} `json:"pi,omitempty"`
}
type Promotion struct {
	ID       []string `json:"id,omitempty"`
	Name     []string `json:"nm,omitempty"`
	Creative []string `json:"cr,omitempty"`
	Position []string `json:"ps,omitempty"`
}

type Client struct {
	ProtocolVersion string `json:"v,omitempty"`
	TrackingID      string `json:"tid,omitempty"`
	AnonymizeIP     string `json:"aip,omitempty"`
	DataSource      string `json:"ds,omitempty"`
	QueueTime       string `json:"qt,omitempty"`
	CacheBuster     string `json:"z,omitempty"`

	ClientID string `json:"cid,omitempty"`
	UserID   string `json:"uid,omitempty"`

	SessionControl       string `json:"sc,omitempty"`
	IPOverride           string `json:"uip,omitempty"`
	UserAgentOverride    string `json:"ua,omitempty"`
	GeographicalOverride string `json:"geoid,omitempty"`

	DocumentReferrer   string `json:"dr,omitempty"`
	CampaignName       string `json:"cn,omitempty"`
	CampaignSource     string `json:"cs,omitempty"`
	CampaignMedium     string `json:"cm,omitempty"`
	CampaignKeyword    string `json:"ck,omitempty"`
	CampaignContent    string `json:"cc,omitempty"`
	CampaignID         string `json:"ci,omitempty"`
	GoogleAdWordsID    string `json:"gclid,omitempty"`
	GoogleDisplayAdsID string `json:"dclid,omitempty"`

	ScreenResolution string `json:"sr,omitempty"`
	ViewportSize     string `json:"vp,omitempty"`
	DocumentEncoding string `json:"de,omitempty"`
	ScreenColors     string `json:"sd,omitempty"`
	UserLanguage     string `json:"ul,omitempty"`
	JavaEnabled      string `json:"je,omitempty"`
	FlashVersion     string `json:"fl,omitempty"`

	HitType           string `json:"t,omitempty"`
	NonInteractionHit string `json:"ni,omitempty"`

	DocumentLocationURL string   `json:"dl,omitempty"`
	DocumentHostName    string   `json:"dh,omitempty"`
	DocumentPath        string   `json:"dp,omitempty"`
	DocumentTitle       string   `json:"dt,omitempty"`
	ScreenName          string   `json:"cd,omitempty"`
	ContentGroup        []string `json:"cg,omitempty"`
	LinkID              string   `json:"linkid,omitempty"`

	ApplicationName        string `json:"an,omitempty"`
	ApplicationID          string `json:"aid,omitempty"`
	ApplicationVersion     string `json:"av,omitempty"`
	ApplicationInstallerID string `json:"aiid,omitempty"`

	EventCategory string `json:"ec,omitempty"`
	EventAction   string `json:"ea,omitempty"`
	EventLabel    string `json:"el,omitempty"`
	EventValue    string `json:"ev,omitempty"`

	TransactionID          string `json:"ti,omitempty"`
	TransactionAffiliation string `json:"ta,omitempty"`
	TransactionRevenue     string `json:"tr,omitempty"`
	TransactionShipping    string `json:"ts,omitempty"`
	TransactionTax         string `json:"tt,omitempty"`
	TransactionCouponCode  string `json:"tcc,omitempty"`

	ItemName     string `json:"in,omitempty"`
	ItemPrice    string `json:"ip,omitempty"`
	ItemQuantity string `json:"iq,omitempty"`
	ItemCode     string `json:"ic,omitempty"`
	ItemCategory string `json:"iv,omitempty"`

	Products           []*Product           `json:"pr,omitempty"`
	ProductImpressions []*ProductImpression `json:"il,omitempty"`
	Promotions         []*Promotion         `json:"promo,omitempty"`
	PromotionAction    string               `json:"promoa,omitempty"`

	CheckoutStep       string `json:"cos,omitempty"`
	CheckoutStepOption string `json:"col,omitempty"`
	CurrencyCode       string `json:"cu,omitempty"`

	SocialNetwork      string `json:"sn,omitempty"`
	SocialAction       string `json:"sa,omitempty"`
	SocialActionTarget string `json:"st,omitempty"`

	UserTimingCategory     string `json:"utc,omitempty"`
	UserTimingVariableName string `json:"utv,omitempty"`
	UserTimingTime         string `json:"utt,omitempty"`
	UserTimingLabel        string `json:"utl,omitempty"`

	PageLoadTime         string `json:"plt,omitempty"`
	PageDownloadTime     string `json:"pdt,omitempty"`
	DNSTime              string `json:"dns,omitempty"`
	RedirectResponseTime string `json:"rrt,omitempty"`
	TCPConnectTime       string `json:"tcp,omitempty"`
	ServerResponseTime   string `json:"srt,omitempty"`
	DOMInteractiveTime   string `json:"dit,omitempty"`
	ContentLoadTime      string `json:"clt,omitempty"`

	ExceptionDescription string `json:"exd,omitempty"`
	IsExceptionFatal     string `json:"exf,omitempty"`

	CustomDimension []string `json:"cd,omitempty"`
	CustomMetric    []string `json:"cm,omitempty"`

	ExperimentID      string `json:"xid,omitempty"`
	ExperimentVariant string `json:"xvar,omitempty"`
}

func (api *API) ParseStruct(prefix string, data map[string]interface{}) (values []string) {
	for k, v := range data {
		if k != "" && v != nil {
			t := reflect.ValueOf(v)
			switch t.Kind() {
			case reflect.Map:
				var iface map[string]interface{}
				r := new(bytes.Buffer)
				encoder := json.NewEncoder(r)
				encoder.Encode(t.Interface())
				decoder := json.NewDecoder(r)
				decoder.UseNumber()
				decoder.Decode(&iface)
				values = append(values, api.ParseStruct(prefix+k, iface)...)
			case reflect.Slice:
				for i := 0; i < t.Len(); i++ {
					f := t.Index(i).Interface()
					ft := reflect.ValueOf(f)
					switch ft.Kind() {
					case reflect.Map:
						var iface map[string]interface{}
						r := new(bytes.Buffer)
						encoder := json.NewEncoder(r)
						encoder.Encode(ft.Interface())
						decoder := json.NewDecoder(r)
						decoder.UseNumber()
						decoder.Decode(&iface)
						values = append(values, api.ParseStruct(prefix+k+fmt.Sprintf("%v", i), iface)...)
					default:
						d := fmt.Sprintf("%v", t)
						if d != "" && d != "0" {
							values = append(values, prefix+k+fmt.Sprintf("%v", i)+"="+d)
						}
					}
				}
			default:
				d := fmt.Sprintf("%v", t)
				if d != "" && d != "0" {
					values = append(values, prefix+k+"="+d)
				}
			}
		}
	}
	return values
}

func (api *API) ParseQuery(str string) string {
	u, _ := url.Parse(str)
	q := u.Query()
	u.RawQuery = q.Encode()
	ret := u.String()
	return strings.TrimLeft(ret, "?")
}

func (api *API) Send(client *Client) string {
	var apidata []string
	var iface map[string]interface{}
	data := new(bytes.Buffer)
	encoder := json.NewEncoder(data)
	encoder.Encode(client)
	decoder := json.NewDecoder(data)
	decoder.UseNumber()
	decoder.Decode(&iface)
	apidata = api.ParseStruct("", iface)
	postdata := api.ParseQuery("?" + strings.Join(apidata, "&"))
	cli := new(http.Client)
	req, err := http.NewRequest("POST", config.ApiUrl, strings.NewReader(postdata))
	if err != nil {
		return err.Error()
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := cli.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	read, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(read)
}
