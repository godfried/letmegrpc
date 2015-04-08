// Code generated by protoc-gen-gogo.
// source: grpc.proto
// DO NOT EDIT!

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	MyRequest
	MyResponse
	MyMsg
	MyMsg2
*/
package grpc

import net_http "net/http"
import encoding_json "encoding/json"
import io "io"
import golang_org_x_net_context "golang.org/x/net/context"
import strings "strings"
import strconv "strconv"
import log "log"
import google_golang_org_grpc "google.golang.org/grpc"

var htmlstringer = func(v interface{}) ([]byte, error) {
	header := []byte("<div class=\"container\">")
	data, err := encoding_json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	footer := []byte("</div>")
	return append(append(header, data...), footer...), nil
}

func SetHtmlStringer(s func(interface{}) ([]byte, error)) {
	htmlstringer = s
}
func Serve(httpAddr, grpcAddr string, opts ...google_golang_org_grpc.DialOption) {
	conn, err := google_golang_org_grpc.Dial(grpcAddr, opts...)
	if err != nil {
		log.Fatalf("Dial(%q) = %v", grpcAddr, err)
	}
	MyTestClient := NewMyTestClient(conn)
	MyTestServer := NewHTMLMyTestServer(MyTestClient)
	net_http.HandleFunc("/MyTest/UnaryCall", MyTestServer.UnaryCall)
	net_http.HandleFunc("/MyTest/Downstream", MyTestServer.Downstream)
	net_http.HandleFunc("/MyTest/Upstream", MyTestServer.Upstream)
	net_http.HandleFunc("/MyTest/Bidi", MyTestServer.Bidi)
	if err := net_http.ListenAndServe(httpAddr, nil); err != nil {
		log.Fatal(err)
	}
}

type htmlMyTest struct {
	client MyTestClient
}

func NewHTMLMyTestServer(client MyTestClient) *htmlMyTest {
	return &htmlMyTest{client}
}
func (this *htmlMyTest) UnaryCall(w net_http.ResponseWriter, req *net_http.Request) {
	w.Write([]byte("<html>"))
	w.Write([]byte("<head>"))
	w.Write([]byte("<title>MyTest - UnaryCall</title>"))
	w.Write([]byte("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css\">"))
	w.Write([]byte("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js\"></script>"))
	w.Write([]byte("<script src=\"//code.jquery.com/jquery-1.11.2.min.js\"></script>"))
	w.Write([]byte("</head>"))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &MyRequest{}
	if len(jsonString) > 0 {
		err := encoding_json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		someValue = true
	} else {
		fieldnames := []string{
			"Value",
			"Value2",
		}
		isString := []bool{
			false,
			false,
		}
		isBool := []bool{
			false,
			false,
		}
		fields := make([]string, 0, len(fieldnames))
		for i, name := range fieldnames {
			v := req.FormValue(name)
			if len(v) > 0 {
				someValue = true
				if isString[i] {
					fields = append(fields, "\""+name+"\":"+strconv.Quote(v))
				} else if isBool[i] {
					if v == "on" {
						fields = append(fields, "\""+name+"\":"+"true")
					} else {
						fields = append(fields, "\""+name+"\":"+"false")
					}
				} else {
					fields = append(fields, "\""+name+"\":"+v)
				}
			}
			if someValue {
				s := "{" + strings.Join(fields, ",") + "}"
				err := encoding_json.Unmarshal([]byte(s), msg)
				if err != nil {
					if err == io.EOF {
						return
					}
					w.Write([]byte(err.Error()))
					return
				}
			}
		}
	}
	w.Write([]byte("<div class=\"container\"><div class=\"jumbotron\">"))
	w.Write([]byte("<h3>MyTest - UnaryCall</h3>"))
	s := "<form action=\"/MyTest/UnaryCall\" method=\"GET\" role=\"form\">"
	w.Write([]byte(s))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value\">Value</label>"))
	w.Write([]byte("<input id=\"Value\" name=\"Value\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value2\">Value2</label>"))
	w.Write([]byte("<input id=\"Value2\" name=\"Value2\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div>"))
	if someValue {
		reply, err := this.client.UnaryCall(golang_org_x_net_context.Background(), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		out, err := htmlstringer(reply)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(out)
	}
	w.Write([]byte("</html>"))
}
func (this *htmlMyTest) Downstream(w net_http.ResponseWriter, req *net_http.Request) {
	w.Write([]byte("<html>"))
	w.Write([]byte("<head>"))
	w.Write([]byte("<title>MyTest - Downstream</title>"))
	w.Write([]byte("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css\">"))
	w.Write([]byte("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js\"></script>"))
	w.Write([]byte("<script src=\"//code.jquery.com/jquery-1.11.2.min.js\"></script>"))
	w.Write([]byte("</head>"))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &MyRequest{}
	if len(jsonString) > 0 {
		err := encoding_json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		someValue = true
	} else {
		fieldnames := []string{
			"Value",
			"Value2",
		}
		isString := []bool{
			false,
			false,
		}
		isBool := []bool{
			false,
			false,
		}
		fields := make([]string, 0, len(fieldnames))
		for i, name := range fieldnames {
			v := req.FormValue(name)
			if len(v) > 0 {
				someValue = true
				if isString[i] {
					fields = append(fields, "\""+name+"\":"+strconv.Quote(v))
				} else if isBool[i] {
					if v == "on" {
						fields = append(fields, "\""+name+"\":"+"true")
					} else {
						fields = append(fields, "\""+name+"\":"+"false")
					}
				} else {
					fields = append(fields, "\""+name+"\":"+v)
				}
			}
			if someValue {
				s := "{" + strings.Join(fields, ",") + "}"
				err := encoding_json.Unmarshal([]byte(s), msg)
				if err != nil {
					if err == io.EOF {
						return
					}
					w.Write([]byte(err.Error()))
					return
				}
			}
		}
	}
	w.Write([]byte("<div class=\"container\"><div class=\"jumbotron\">"))
	w.Write([]byte("<h3>MyTest - Downstream</h3>"))
	s := "<form action=\"/MyTest/Downstream\" method=\"GET\" role=\"form\">"
	w.Write([]byte(s))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value\">Value</label>"))
	w.Write([]byte("<input id=\"Value\" name=\"Value\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value2\">Value2</label>"))
	w.Write([]byte("<input id=\"Value2\" name=\"Value2\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div>"))
	if someValue {
		down, err := this.client.Downstream(golang_org_x_net_context.Background(), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		for {
			reply, err := down.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				w.Write([]byte(err.Error()))
				return
			}
			out, err := htmlstringer(reply)
			if err != nil {
				if err == io.EOF {
					return
				}
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte("<p>"))
			w.Write(out)
			w.Write([]byte("</p>"))
			w.(net_http.Flusher).Flush()
		}
	}
	w.Write([]byte("</html>"))
}
func (this *htmlMyTest) Upstream(w net_http.ResponseWriter, req *net_http.Request) {
	w.Write([]byte("<html>"))
	w.Write([]byte("<head>"))
	w.Write([]byte("<title>MyTest - Upstream</title>"))
	w.Write([]byte("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css\">"))
	w.Write([]byte("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js\"></script>"))
	w.Write([]byte("<script src=\"//code.jquery.com/jquery-1.11.2.min.js\"></script>"))
	w.Write([]byte("</head>"))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &MyMsg{}
	if len(jsonString) > 0 {
		err := encoding_json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		someValue = true
	} else {
		fieldnames := []string{
			"Value",
		}
		isString := []bool{
			false,
		}
		isBool := []bool{
			false,
		}
		fields := make([]string, 0, len(fieldnames))
		for i, name := range fieldnames {
			v := req.FormValue(name)
			if len(v) > 0 {
				someValue = true
				if isString[i] {
					fields = append(fields, "\""+name+"\":"+strconv.Quote(v))
				} else if isBool[i] {
					if v == "on" {
						fields = append(fields, "\""+name+"\":"+"true")
					} else {
						fields = append(fields, "\""+name+"\":"+"false")
					}
				} else {
					fields = append(fields, "\""+name+"\":"+v)
				}
			}
			if someValue {
				s := "{" + strings.Join(fields, ",") + "}"
				err := encoding_json.Unmarshal([]byte(s), msg)
				if err != nil {
					if err == io.EOF {
						return
					}
					w.Write([]byte(err.Error()))
					return
				}
			}
		}
	}
	w.Write([]byte("<div class=\"container\"><div class=\"jumbotron\">"))
	w.Write([]byte("<h3>MyTest - Upstream</h3>"))
	s := "<form action=\"/MyTest/Upstream\" method=\"GET\" role=\"form\">"
	w.Write([]byte(s))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value\">Value</label>"))
	w.Write([]byte("<input id=\"Value\" name=\"Value\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div>"))
	if someValue {
		up, err := this.client.Upstream(golang_org_x_net_context.Background())
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		err = up.Send(msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		reply, err := up.CloseAndRecv()
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		out, err := htmlstringer(reply)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(out)
	}
	w.Write([]byte("</html>"))
}
func (this *htmlMyTest) Bidi(w net_http.ResponseWriter, req *net_http.Request) {
	w.Write([]byte("<html>"))
	w.Write([]byte("<head>"))
	w.Write([]byte("<title>MyTest - Bidi</title>"))
	w.Write([]byte("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css\">"))
	w.Write([]byte("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js\"></script>"))
	w.Write([]byte("<script src=\"//code.jquery.com/jquery-1.11.2.min.js\"></script>"))
	w.Write([]byte("</head>"))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &MyMsg{}
	if len(jsonString) > 0 {
		err := encoding_json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		someValue = true
	} else {
		fieldnames := []string{
			"Value",
		}
		isString := []bool{
			false,
		}
		isBool := []bool{
			false,
		}
		fields := make([]string, 0, len(fieldnames))
		for i, name := range fieldnames {
			v := req.FormValue(name)
			if len(v) > 0 {
				someValue = true
				if isString[i] {
					fields = append(fields, "\""+name+"\":"+strconv.Quote(v))
				} else if isBool[i] {
					if v == "on" {
						fields = append(fields, "\""+name+"\":"+"true")
					} else {
						fields = append(fields, "\""+name+"\":"+"false")
					}
				} else {
					fields = append(fields, "\""+name+"\":"+v)
				}
			}
			if someValue {
				s := "{" + strings.Join(fields, ",") + "}"
				err := encoding_json.Unmarshal([]byte(s), msg)
				if err != nil {
					if err == io.EOF {
						return
					}
					w.Write([]byte(err.Error()))
					return
				}
			}
		}
	}
	w.Write([]byte("<div class=\"container\"><div class=\"jumbotron\">"))
	w.Write([]byte("<h3>MyTest - Bidi</h3>"))
	s := "<form action=\"/MyTest/Bidi\" method=\"GET\" role=\"form\">"
	w.Write([]byte(s))
	w.Write([]byte("<div class=\"form-group\">"))
	w.Write([]byte("<label for=\"Value\">Value</label>"))
	w.Write([]byte("<input id=\"Value\" name=\"Value\" type=\"text\" class=\"form-control\"/><br>"))
	w.Write([]byte("</div>"))
	w.Write([]byte("<button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div>"))
	if someValue {
		bidi, err := this.client.Bidi(golang_org_x_net_context.Background())
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		err = bidi.Send(msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		reply, err := bidi.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		out, err := htmlstringer(reply)
		if err != nil {
			if err == io.EOF {
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(out)
	}
	w.Write([]byte("</html>"))
}
