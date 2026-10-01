---
page_title: "discovery_k8s.access_info.kubeconfig_url.clear_secret_info"
subcategory: ""
description: "discovery_k8s.access_info.kubeconfig_url.clear_secret_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 3632, "body_sha256": "sha256:af5504fae7326d2b797fca06bfeff130ae5ff512929741244016aeb7b12a7771", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:clear_secret_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "path": "docs/guides/resources--discovery--properties--discovery_k8s--access_info--kubeconfig_url--clear_secret_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "access_info", "kubeconfig_url", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.access_info.kubeconfig_url.clear_secret_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info.kubeconfig_url.clear_secret_info

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md)
- [discovery_k8s.access_info](resources--discovery--properties--discovery_k8s--access_info.md)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--properties--discovery_k8s--access_info--kubeconfig_url.md)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--properties--discovery_k8s--access_info--kubeconfig_url.md)
- [xcsh_discovery](../resources/discovery.md)
