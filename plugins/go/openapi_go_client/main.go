
package main

import "github.com/googleapis/openapi-compiler/plugins/go/template_plugin"

func main() {
    template_plugin.Run(map[string]string{ 
        "client.go": "cGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCiAgImVuY29kaW5nL2pzb24iCiAgImZtdCIKICAibmV0L2h0dHAiCiAgInN0cmluZ3MiCikKICAKY29uc3Qgc2VydmljZSA9ICJodHRwOi8vbG9jYWxob3N0OjgwODAiCgp7e3JhbmdlIC5SZW5kZXJlci5NZXRob2RzfX0Ke3tpZiBlcSAuUmVzdWx0VHlwZU5hbWUgIiJ9fQpmdW5jIHt7LkNsaWVudE5hbWV9fSh7e3BhcmFtZXRlckxpc3QgLn19KSAoZXJyIGVycm9yKSB7Cnt7ZWxzZX19CmZ1bmMge3suQ2xpZW50TmFtZX19KHt7cGFyYW1ldGVyTGlzdCAufX0pIChyZXN1bHQgKnt7LlJlc3VsdFR5cGVOYW1lfX0sIGVyciBlcnJvcikgewp7e2VuZH19CglwYXRoIDo9IHNlcnZpY2UgKyAie3suUGF0aH19IgoJCgl7e3JhbmdlIC5QYXJhbWV0ZXJzVHlwZS5GaWVsZHN9fQkKCXt7aWYgZXEgLlBvc2l0aW9uICJwYXRoIn19CglwYXRoID0gc3RyaW5ncy5SZXBsYWNlKHBhdGgsICJ7IiArICJ7ey5KU09OTmFtZX19IiArICJ9IiwgZm10LlNwcmludGYoIiV2Iiwge3suSlNPTk5hbWV9fSksIDEpCgl7e2VuZH19Cgl7e2VuZH19CgkKCXJlcSwgZXJyIDo9IGh0dHAuTmV3UmVxdWVzdCgie3suTWV0aG9kfX0iLCBwYXRoLCBuaWwpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4KCX0KCXJlc3AsIGVyciA6PSBodHRwLkRlZmF1bHRDbGllbnQuRG8ocmVxKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIAoJfQoJZGVmZXIgcmVzcC5Cb2R5LkNsb3NlKCkKCXt7aWYgbmUgLlJlc3VsdFR5cGVOYW1lICIifX0KCWRlY29kZXIgOj0ganNvbi5OZXdEZWNvZGVyKHJlc3AuQm9keSkKCXJlc3VsdCA9ICZ7ey5SZXN1bHRUeXBlTmFtZX19e30KCWRlY29kZXIuRGVjb2RlKHJlc3VsdCkKCXt7ZW5kfX0KCXJldHVybiAgCn0Ke3tlbmR9fQo=",
        "types.go": "cGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCikKCi8vIHR5cGVzCnt7cmFuZ2UgLlJlbmRlcmVyLlR5cGVzfX0KdHlwZSB7ey5OYW1lfX0gc3RydWN0IHsge3tyYW5nZSAuRmllbGRzfX0KICB7ey5OYW1lfX0ge3suVHlwZX19e3tpZiBuZSAuSlNPTk5hbWUgIiJ9fSBganNvbjoie3suSlNPTk5hbWV9fSJge3tlbmR9fXt7ZW5kfX0KfQp7e2VuZH19Cg==",
    })
}