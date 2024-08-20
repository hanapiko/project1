package models

type DateEntry struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}

type DateResponse struct {
    Index []DateEntry `json:"index"`
}
