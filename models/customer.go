ackage models
import (
    "fmt"
    "net/http"
)
type Item struct {
    ID int `json:"id"`
    customer int `json:"customer"`
    services int `json: "services"`   
    orders int `json: "orders"`
    balance float32  `json: "balance"`
}
type ItemList struct {
    Items []Item `json:"items"`
}
func (i *Item) Bind(r *http.Request) error {
    if i.customer == "" {
        return fmt.Errorf("customer is a required field")
    }
    return nil
}
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}