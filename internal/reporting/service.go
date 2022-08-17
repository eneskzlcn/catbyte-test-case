package reporting

type RedisClient interface {
	Get(key string, out interface{}) error
}
type Service struct {
	client RedisClient
}

func NewService(client RedisClient) *Service {
	return &Service{client: client}
}
func (s *Service) Report(sender, reciever string) [][]Report {
	//s.client.Get("")
	// Not enough time to complete.
	return nil
}
