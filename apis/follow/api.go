package follow

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/rpc"

	// Vendor
	"github.com/pkg/errors"
)

const APIID = "follow_api"

type API struct {
	id     int
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) (*API, error) {
	id, err := rpc.GetNumericAPIID(caller, APIID)
	if err != nil {
		return nil, err
	}
	return &API{id, caller}, nil
}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{api.id, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

func (api *API) GetFollowersRaw(accountName, start, kind string, limit uint16) (*json.RawMessage, error) {
	return api.Raw("get_followers", []interface{}{accountName, start, kind, limit})
}

func (api *API) GetFollowers(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	raw, err := api.GetFollowersRaw(accountName, start, kind, limit)
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_followers response")
	}
	return resp, nil
}

func (api *API) GetFollowingRaw(accountName, start, kind string, limit uint16) (*json.RawMessage, error) {
	return api.Raw("get_following", []interface{}{accountName, start, kind, limit})
}

func (api *API) GetFollowing(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	raw, err := api.GetFollowingRaw(accountName, start, kind, limit)
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_following response")
	}
	return resp, nil
}

func (api *API) GetFollowCountRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_follow_count", []interface{}{accountName})
}

func (api *API) GetFollowCount(accountName string) (*FollowCount, error) {
	raw, err := api.GetFollowCountRaw(accountName)
	if err != nil {
		return nil, err
	}
	var resp *FollowCount
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_follow_count response")
	}
	return resp, nil
}

func (api *API) GetFeedEntriesRaw(accountName string, entryID uint32, limit uint16) (*json.RawMessage, error) {
	if limit > 500 {
		return nil, errors.New("golos-go: follow_api: get_feed_entries -> limit must not exceed 500")
	}
	return api.Raw("get_feed_entries", []interface{}{accountName, entryID, limit})
}

func (api *API) GetFeedEntries(accountName string, entryID uint32, limit uint16) ([]*FeedEntry, error) {
	raw, err := api.GetFeedEntriesRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}
	var resp []*FeedEntry
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

func (api *API) GetFeedRaw(accountName string, entryID uint32, limit uint16) (*json.RawMessage, error) {
	if limit > 500 {
		return nil, errors.New("golos-go: follow_api: get_feed -> limit must not exceed 500")
	}
	return api.Raw("get_feed", []interface{}{accountName, entryID, limit})
}

func (api *API) GetFeed(accountName string, entryID uint32, limit uint16) ([]*Feeds, error) {
	raw, err := api.GetFeedRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Feeds
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_feed response")
	}
	return resp, nil
}

func (api *API) GetBlogEntriesRaw(accountName string, entryID uint32, limit uint16) (*json.RawMessage, error) {
	if limit > 500 {
		return nil, errors.New("golos-go: follow_api: get_blog_entries -> limit must not exceed 500")
	}
	return api.Raw("get_blog_entries", []interface{}{accountName, entryID, limit})
}

func (api *API) GetBlogEntries(accountName string, entryID uint32, limit uint16) ([]*BlogEntries, error) {
	raw, err := api.GetBlogEntriesRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}
	var resp []*BlogEntries
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

func (api *API) GetBlogRaw(accountName string, entryID uint32, limit uint16) (*json.RawMessage, error) {
	if limit > 500 {
		return nil, errors.New("golos-go: follow_api: get_blog -> limit must not exceed 500")
	}
	return api.Raw("get_blog", []interface{}{accountName, entryID, limit})
}

func (api *API) GetBlog(accountName string, entryID uint32, limit uint16) ([]*Blogs, error) {
	raw, err := api.GetBlogRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Blogs
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_feed response")
	}
	return resp, nil
}

func (api *API) GetAccountReputationsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("golos-go: follow_api: get_account_reputations -> limit must not exceed 1000")
	}
	return api.Raw("get_account_reputations", []interface{}{lowerBoundName, limit})
}

func (api *API) GetAccountReputations(lowerBoundName string, limit uint32) ([]*AccountReputation, error) {
	raw, err := api.GetAccountReputationsRaw(lowerBoundName, limit)
	if err != nil {
		return nil, err
	}
	var resp []*AccountReputation
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: follow_api: failed to unmarshal get_account_reputations response")
	}
	return resp, nil
}

func (api *API) GetRebloggedByRaw(author, permlink string) (*json.RawMessage, error) {
	return api.Raw("get_reblogged_by", []interface{}{author, permlink})
}

func (api *API) GetRebloggedBy(author, permlink string) ([]string, error) {
	raw, err := api.GetRebloggedByRaw(author, permlink)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_reblogged_by response")
	}
	return resp, nil
}

func (api *API) GetBlogAuthorsRaw(author string) (*json.RawMessage, error) {
	return api.Raw("get_blog_authors", []interface{}{author})
}

func (api *API) GetBlogAuthors(author string) (*BlogAuthors, error) {
	raw, err := api.GetBlogAuthorsRaw(author)
	if err != nil {
		return nil, err
	}
	var resp BlogAuthors
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_blog_authors response")
	}
	return &resp, nil
}
