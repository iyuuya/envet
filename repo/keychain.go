package repo

import (
	"fmt"
	"strings"

	keychain "github.com/keybase/go-keychain"
)

type Entry struct {
	Namespace string
	Key string
	Value string
}

type KeychainRepository struct {}

func NewKeychainRepository() *KeychainRepository {
	return &KeychainRepository{}
}

func (r *KeychainRepository) ListAll() ([]Entry, error) {
	var entries []Entry

	q := keychain.NewItem()
	q.SetSecClass(keychain.SecClassGenericPassword)
	q.SetMatchLimit(keychain.MatchLimitAll)
	q.SetReturnAttributes(true)

	attrs, err := keychain.QueryItem(q)
	if err != nil {
		return nil, err
	}

	for _, it := range attrs {
		if !strings.HasPrefix(it.Service, "envchain-") {
			continue
		}
		qv := keychain.NewItem()
		qv.SetSecClass(keychain.SecClassGenericPassword)
		qv.SetService(it.Service)
		qv.SetAccount(it.Account)
		qv.SetReturnData(true)
		qv.SetMatchLimit(keychain.MatchLimitOne)

		val, err := keychain.QueryItem(qv)
		if err != nil {
			return nil, fmt.Errorf("ERROR %s/%s: %v", it.Service, it.Account, err)
		}
		if len(val) == 0 {
			continue
		}

		entry := Entry{
			Namespace: strings.TrimPrefix(it.Service, "envchain-"),
			Key:   it.Account,
			Value: string(val[0].Data),
		}

		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *KeychainRepository) ListNamespaces() ([]string, error) {
	var namespaces []string

	q := keychain.NewItem()
	q.SetSecClass(keychain.SecClassGenericPassword)
	q.SetMatchLimit(keychain.MatchLimitAll)
	q.SetReturnAttributes(true)

	attrs, err := keychain.QueryItem(q)
	if err != nil {
		return nil, err
	}

	nsMap := make(map[string]struct{})
	for _, it := range attrs {
		if strings.HasPrefix(it.Service, "envchain-") {
			ns, _ := strings.CutPrefix(it.Service, "envchain-")
			nsMap[ns] = struct{}{}
		}
	}

	for ns := range nsMap {
		namespaces = append(namespaces, ns)
	}

	return namespaces, nil
}

func (r *KeychainRepository) ListKeys(namespace string) ([]string, error) {
	var entries []string

	q := keychain.NewItem()
	q.SetSecClass(keychain.SecClassGenericPassword)
	q.SetService("envchain-" + namespace)
	q.SetMatchLimit(keychain.MatchLimitAll)
	q.SetReturnAttributes(true)

	attrs, err := keychain.QueryItem(q)
	if err != nil {
		return nil, err
	}

	for _, it := range attrs {
		entries = append(entries, it.Account)
	}

	return entries, nil
}

func (r *KeychainRepository) ListKeyValues(namespace string) ([]Entry, error) {
	var entries []Entry

	q := keychain.NewItem()
	q.SetSecClass(keychain.SecClassGenericPassword)
	q.SetService("envchain-" + namespace)
	q.SetMatchLimit(keychain.MatchLimitAll)
	q.SetReturnAttributes(true)
	attrs, err := keychain.QueryItem(q)
	if err != nil {
		return nil, err
	}

	for _, it := range attrs {
		qv := keychain.NewItem()
		qv.SetSecClass(keychain.SecClassGenericPassword)
		qv.SetService(it.Service)
		qv.SetAccount(it.Account)
		qv.SetReturnData(true)
		qv.SetMatchLimit(keychain.MatchLimitOne)
		val, err := keychain.QueryItem(qv)
		if err != nil {
			return nil, fmt.Errorf("ERROR %s/%s: %v", it.Service, it.Account, err)
		}
		if len(val) == 0 {
			continue
		}
		entry := Entry{
			Namespace: strings.TrimPrefix(it.Service, "envchain-"),
			Key:   it.Account,
			Value: string(val[0].Data),
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *KeychainRepository) AddEntry(namespace, key, value string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService("envchain-" + namespace)
	item.SetAccount(key)
	item.SetData([]byte(value))
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	err := keychain.AddItem(item)
	return err
}

func (r *KeychainRepository) RemoveEntry(namespace, key string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService("envchain-" + namespace)
	item.SetAccount(key)

	err := keychain.DeleteItem(item)
	return err
}

func (r *KeychainRepository) RemoveNamespace(namespace string) error {
	q := keychain.NewItem()
	q.SetSecClass(keychain.SecClassGenericPassword)
	q.SetService("envchain-" + namespace)
	q.SetMatchLimit(keychain.MatchLimitAll)
	q.SetReturnAttributes(true)

	attrs, err := keychain.QueryItem(q)
	if err != nil {
		return err
	}

	for _, it := range attrs {
		item := keychain.NewItem()
		item.SetSecClass(keychain.SecClassGenericPassword)
		item.SetService(it.Service)
		item.SetAccount(it.Account)

		err := keychain.DeleteItem(item)
		if err != nil {
			return err
		}
	}

	return nil
}
