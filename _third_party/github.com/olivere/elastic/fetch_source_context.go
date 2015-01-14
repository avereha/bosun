package elastic

type FetchSourceContext struct {
	fetchSource     bool
	transformSource bool
	includes        []string
	excludes        []string
}

func NewFetchSourceContext(fetchSource bool) *FetchSourceContext {
	return &FetchSourceContext{
		fetchSource: fetchSource,
		includes:    make([]string, 0),
		excludes:    make([]string, 0),
	}
}

func (fsc *FetchSourceContext) FetchSource() bool {
	return fsc.fetchSource
}

func (fsc *FetchSourceContext) SetFetchSource(fetchSource bool) {
	fsc.fetchSource = fetchSource
}

func (fsc *FetchSourceContext) Include(includes ...string) *FetchSourceContext {
	fsc.includes = append(fsc.includes, includes...)
	return fsc
}

func (fsc *FetchSourceContext) Exclude(excludes ...string) *FetchSourceContext {
	fsc.excludes = append(fsc.excludes, excludes...)
	return fsc
}

func (fsc *FetchSourceContext) TransformSource(transformSource bool) *FetchSourceContext {
	fsc.transformSource = transformSource
	return fsc
}

func (fsc *FetchSourceContext) Source() interface{} {
	if !fsc.fetchSource {
		return false
	}
	return map[string]interface{}{
		"includes": fsc.includes,
		"excludes": fsc.excludes,
	}
}
