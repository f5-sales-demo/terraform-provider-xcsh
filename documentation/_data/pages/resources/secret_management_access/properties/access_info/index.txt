---
page_title: "access_info"
subcategory: ""
description: "HostAccessInfoType contains the information about how to connect to the remote host."
xcsh_docs: {"aliases": ["access info"], "body_bytes": 2974, "body_sha256": "sha256:c4de3401c94258c7f5d4239b968812d32793c0c1348a5c74624efc3464477aa5", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info", "parent_id": "xcsh-docs:resources:secret_management_access:reference", "path": "documentation/resources/secret_management_access/properties/access_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info"], "schema_version": 1, "sections": [{"aliases": ["access info rest auth info"], "anchor": "section", "description": "Authentication parameters for REST based hosts.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info"], "syntax": "block", "type": "object"}, {"aliases": ["access info scheme"], "anchor": "schema-access_info--scheme", "description": "SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "scheme"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info server endpoint"], "anchor": "schema-access_info--server_endpoint", "description": "Endpoint to connect to, in host:port format.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "server_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config"], "anchor": "section", "description": "TLS configuration for upstream connections.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config"], "syntax": "block", "type": "object"}, {"aliases": ["access info vault auth info"], "anchor": "section", "description": "Authentication parameters for Hashicorp Vault hosts.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "HostAccessInfoType contains the information about how to connect to the remote host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HostAccessInfoType contains the information about how to connect to the remote host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/): complete subsection reference.

<a id="schema-access_info--scheme"></a>

### scheme property

Type: `"string"`. Optional.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "HTTP",
  "enum": [
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--server_endpoint"></a>

### server_endpoint property

Type: `"string"`. Optional.

Endpoint to connect to, in host:port format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/): complete subsection reference.

- [vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/): complete subsection reference.
