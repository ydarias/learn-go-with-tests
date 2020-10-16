package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrKeyNotFound             = DictionaryErr("key cannot be found at the dictionary")
	ErrKeyAlreadyExists        = DictionaryErr("the given key already exists at the dictionary")
	ErrKeyToUpdateDoesNotExist = DictionaryErr("the given key doesn't exist yet at the dictionary")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, content string) error {
	_, ok := d[key]

	if ok {
		return ErrKeyAlreadyExists
	}

	d[key] = content
	return nil
}

func (d Dictionary) Update(key, content string) error {
	_, ok := d[key]

	if !ok {
		return ErrKeyToUpdateDoesNotExist
	}

	d[key] = content
	return nil
}
