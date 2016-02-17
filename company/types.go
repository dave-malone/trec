package trec

type companyRepository interface {
	addCompany(company Company) (err error)
	getCompanies() (companies []Company)
	getCompany(id string) (company Company, err error)
}

type Company struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
