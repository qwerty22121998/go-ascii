package server

func (s *Server) Register() {
	s.e.GET("/url", s.controller.FromURL.ToText)
	s.e.GET("/url/image", s.controller.FromURL.ToImage)
}
