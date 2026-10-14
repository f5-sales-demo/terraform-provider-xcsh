// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package mocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) handleSwagger(w http.ResponseWriter, r *http.Request) bool {
	prefix := "/api/object_store/namespaces/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		return false
	}
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, prefix), "/")
	if len(parts) < 3 || parts[1] != "stored_objects" || parts[2] != "swagger" {
		return false
	}
	ns := parts[0]
	switch {
	case r.Method == "GET" && len(parts) == 3:
		name := r.URL.Query().Get("name")
		items := []interface{}{}
		versions := []interface{}{}
		s.mu.RLock()
		for path := range s.resources {
			expected := r.URL.Path + "/" + name + "/"
			if strings.HasPrefix(path, expected) {
				versions = append(versions, map[string]string{"version": strings.TrimPrefix(path, expected)})
			}
		}
		s.mu.RUnlock()
		if len(versions) > 0 {
			items = append(items, map[string]interface{}{"name": name, "versions": versions})
		}
		s.writeJSONResponse(w, 200, map[string]interface{}{"items": items})
	case r.Method == "PUT" && len(parts) == 4:
		var request struct {
			Namespace string `json:"namespace"`
			Name      string `json:"name"`
			Content   string `json:"string_value"`
		}
		if json.NewDecoder(r.Body).Decode(&request) != nil || request.Namespace != ns || request.Name != parts[3] {
			s.writeErrorResponse(w, 400, "BAD_REQUEST", "invalid Swagger identity")
			return true
		}
		s.mu.Lock()
		s.resourceVersion++
		version := fmt.Sprintf("v%d", s.resourceVersion)
		object := map[string]interface{}{"metadata": map[string]string{"namespace": ns, "name": request.Name, "version": version}, "string_value": request.Content, "content_format": "json"}
		s.resources[r.URL.Path+"/"+version] = object
		s.mu.Unlock()
		s.writeJSONResponse(w, 200, object)
	case r.Method == "GET" && len(parts) == 5:
		s.handleGet(w, r, r.URL.Path)
	case r.Method == "DELETE" && len(parts) == 5:
		s.handleDelete(w, r, r.URL.Path)
	default:
		s.writeErrorResponse(w, 405, "METHOD_NOT_ALLOWED", "invalid Swagger operation")
	}
	return true
}
