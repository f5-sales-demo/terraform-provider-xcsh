---
page_title: "code_base_integration.azure_repos.access_token.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["code base integration azure repos access token clear secret info"], "body_bytes": 4174, "body_sha256": "sha256:77be40a01c6df4d6ed386fa5026dc86ace4fec116a6406a626daf196104e6e34", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos:access_token:clear_secret_info", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos:access_token", "path": "documentation/resources/code_base_integration/properties/code_base_integration/azure_repos/access_token/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3323010223310032-1331233001013122-0202123111302331-1330312320022321-3312231032331001-2101331312221322-1212130320102233-2112010031031110", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "schema-code_base_integration--azure_repos--access_token--clear_secret_info--url", "enforcement": "provider-schema", "group": "code_base_integration.azure_repos.access_token.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos:access_token:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "azure_repos", "access_token", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-code_base_integration--azure_repos--access_token--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos:access_token:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "azure_repos", "access_token", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-code_base_integration--azure_repos--access_token--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos:access_token:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "azure_repos", "access_token", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/azure_repos/access_token/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.azure_repos.access_token.clear_secret_info

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/azure_repos/)
- [code_base_integration.azure_repos.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/azure_repos/access_token/)
- code_base_integration.azure_repos.access_token.clear_secret_info

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

<a id="schema-code_base_integration--azure_repos--access_token--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-code_base_integration--azure_repos--access_token--clear_secret_info--url"></a>

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

- [code_base_integration.azure_repos.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/azure_repos/access_token/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
