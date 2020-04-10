package models

type Sub struct{
	Adsh       *string    `csv:"adsh"`
	Cik        *string    `csv:"cik"`
	Name       *string    `csv:"name"`
	Sic        *string    `csv:"-"`
	Countryba  *string    `csv:"-"`
	Stprba     *string    `csv:"-"`
	Cityba     *string    `csv:"-"`
	Zipba      *string    `csv:"-"`
	Bas1       *string    `csv:"-"`
	Bas2       *string    `csv:"-"`
	Baph       *string    `csv:"-"`
	Countryma  *string    `csv:"-"`
	Stprma     *string    `csv:"-"`
	Cityma     *string    `csv:"-"`
	Zipma      *string    `csv:"-"`
	Mas1       *string    `csv:"-"`
	Mas2       *string    `csv:"-"`
	Countryinc *string    `csv:"-"`
	Stprinc    *string    `csv:"-"`
	Ein        *string    `csv:"-"`
	Former     *string    `csv:"former"`
	Changed    *string    `csv:"changed"`
	Afs        *string    `csv:"afs"`
	Wksi       *bool      `csv:"wksi"`
	Fye        *string    `csv:"fye"`
	Form       FormType   `csv:"form"`
	Period     Date    `csv:"period"`
	Fy         Year      `csv:"fy"`
	Fp         *string       `csv:"fp"` // TODO probably needs to be FPF
	Filed      Date    `csv:"filed"`
	Accepted   *string    `csv:"accepted"`
	Prevrpt    bool      `csv:"prevrpt"`
	Detail     bool      `csv:"detail"`
	Instance   *string    `csv:"instance"`
	Nciks      *int       `csv:"nciks"`
	Aciks      *string    `csv:"aciks"`
}
