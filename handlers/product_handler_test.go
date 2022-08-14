package handlers_test

import (
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"kube-go-api-tutorial/db/models"
	"kube-go-api-tutorial/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProductHandler(t *testing.T){
	//setup request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	w:= httptest.NewRecorder()

	testObj := new(MyMockedObject)
	testObj.On("WriteToDb", &models.Product{Code: "DC", Price: 100}).Return(nil)

	handlers.ProductHandler{DbWriter: testObj}.CreateProductHandler(w, req)

	res:= w.Result()
	defer res.Body.Close()
	_, err:= ioutil.ReadAll(res.Body)
	if err!=nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

func TestGetProductRequestHandler(t *testing.T) {
	// setup request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("GetFirstByCode", "DC").Return(models.Product{}, nil)

	//set the handler for http request
	handlers.ProductHandler{DbReader: testObj}.GetProductRequestHandler(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
}

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) GetFirstByCode(code string) (models.Product, error) {
	args := m.Called(code)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MyMockedObject) WriteToDb(product *models.Product) {}
