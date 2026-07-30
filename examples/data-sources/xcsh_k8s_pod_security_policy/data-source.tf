# K8SPodSecurityPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityPolicy by name
data "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}

output "k8s_pod_security_policy_id" {
  value = data.xcsh_k8s_pod_security_policy.example.id
}
