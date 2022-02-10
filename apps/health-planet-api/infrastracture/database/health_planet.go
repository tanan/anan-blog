package database

import (
	//"health-planet-api/infra/database/model"
)

func (h SQLHandler) FindByID(id domain.ContractID) (contract domain.Contract, err error) {
	var m model.Contract
	db := h.Conn.Where("id = UUID_TO_BIN(?)", id).First(&m)
	if db.Error != nil && !db.RecordNotFound() {
		return domain.Contract{}, db.Error
	}
	contract = h.toContract(m)

	contractor, err := h.FindContractorByContractID(id)
	if err != nil {
		return domain.Contract{}, err
	}
	editor, err := h.FindEditorByContractID(id)
	categories, err := h.FindCategoriesByContractID(id)
	if err != nil {
		return domain.Contract{}, err
	}
	contract.Categories = categories
	contract.Contractor = contractor
	contract.Editor = editor
	return contract, err
}

func (h SQLHandler) FindCategoriesByContractID(id domain.ContractID) (categories domain.ContractCategories, err error) {
	var m []model.Category
	db := h.Conn.
		Joins("left join contract_category on contract_category.category_id = category.id").
		Where("contract_category.id = UUID_TO_BIN(?)", id).
		Find(&m)
	if db.Error != nil && !db.RecordNotFound() {
		return nil, db.Error
	}
	return h.toCategory(m), err
}

func (h SQLHandler) FindContractorByContractID(id domain.ContractID) (contractor domain.Contractor, err error) {
	var m model.Contractor
	db := h.Conn.
		Joins("left join contract_contractor on contract_contractor.id = contractor.id").
		Where("contract_contractor.id = UUID_TO_BIN(?)", id).
		First(&m)
	if db.Error != nil && !db.RecordNotFound() {
		return domain.Contractor{}, db.Error
	}
	return h.toContractor(m), nil
}

func (h SQLHandler) FindEditorByContractID(id domain.ContractID) (editor domain.Editor, err error) {
	var m model.Editor
	db := h.Conn.
		Joins("left join contract_editor on contract_editor.id = editor.id").
		Where("contract_editor.id = UUID_TO_BIN(?)", id).
		First(&m)
	if db.Error != nil && !db.RecordNotFound() {
		return domain.Editor{}, db.Error
	}
	return h.toEditor(m), nil
}

func (h SQLHandler) toContract(contract model.Contract) domain.Contract {
	return domain.Contract{
		ID:          domain.ContractID(contract.ID.String()),
		Title:       contract.Name,
		Description: contract.Description,
		Note:        contract.Note,
		ContractPeriod: domain.ContractPeriod{
			StartDate: contract.StartDate,
			EndDate:   contract.EndDate,
		},
		RegistrationTime: contract.RegistrationTime,
		UpdateTime:       contract.UpdateTime,
	}
}

func (h SQLHandler) toContractor(contractor model.Contractor) domain.Contractor {
	return domain.Contractor{
		ID:   domain.ContractorID(contractor.ID.String()),
		Name: contractor.Name,
	}
}

func (h SQLHandler) toEditor(contractor model.Editor) domain.Editor {
	return domain.Editor{
		ID:   domain.EditorID(contractor.ID.String()),
		Name: contractor.Name,
	}
}

func (h SQLHandler) toCategory(category []model.Category) (categories domain.ContractCategories) {
	for _, v := range category {
		categories = append(categories, domain.ContractCategory{
			ID:   domain.ContractCategoryID(v.ID),
			Name: v.Name,
		})
	}
	return categories
}

func (h SQLHandler) CreateContract(contract domain.Contract) (c domain.Contract, err error) {
	m := model.Contract{
		Name:             contract.Title,
		Description:      contract.Description,
		Note:             contract.Note,
		RegistrationTime: contract.RegistrationTime,
		UpdateTime:       contract.UpdateTime,
		StartDate:        contract.ContractPeriod.StartDate,
		EndDate:          contract.ContractPeriod.EndDate,
	}
	if db := h.Conn.Create(&m); db.Error != nil {
		return domain.Contract{}, db.Error
	}
	contract.ID = domain.ContractID(m.ID.String())
	return contract, nil
}