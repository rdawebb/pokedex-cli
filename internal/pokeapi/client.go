package pokeapi

type Client struct {
	baseUrl    string
	nextUrl    string
	previousUrl string
}

func NewClient() *Client {
	return &Client{
		baseUrl: "https://pokeapi.co/api/v2/location-area?limit=20",
	}
}

func (client *Client) Reset() {
		client.nextUrl = ""
		client.previousUrl = ""
}

func (client *Client) GetNextUrl() string {
	return client.nextUrl
}

func (client *Client) GetPreviousUrl() string {
	return client.previousUrl
}

func (client *Client) GetBaseUrl() string {
	return client.baseUrl
}
