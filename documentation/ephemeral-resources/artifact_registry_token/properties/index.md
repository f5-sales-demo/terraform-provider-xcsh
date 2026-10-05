---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_artifact_registry_token."
xcsh_docs: {"aliases": ["artifact registry token"], "body_bytes": 1705, "body_sha256": "sha256:02adc11374d708ee63882384d92540be109e9e5393f7eb28c1bf14bb0e9c0f44", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "parent_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:fundamentals", "path": "documentation/ephemeral-resources/artifact_registry_token/properties/index.md", "product": "distributed-cloud", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-2203313033322030-0113022202301111-0210200113320313-3110001101012332-3113033331112210-1332120310221123-3101023200110313-3312021311022023", "registry_path": "docs/guides/ephemeral-resources--artifact_registry_token--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["expiration time"], "anchor": "schema-expiration_time", "description": "Expiration Time. Expiration time of the token.", "document_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expiration_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace is used to scope the query.", "document_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["token"], "anchor": "schema-token", "description": "Access token for the F5 Artifact Registry (FAR) This token can be used to authenticate with FAR when pulling related images for Kubernetes bot infrastructure.", "document_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["token"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_artifact_registry_token.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_artifact_registry_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/)
- Property reference

## Direct properties

<a id="schema-expiration_time"></a>

### expiration_time property

Type: `"string"`. Computed, Sensitive.

Expiration Time. Expiration time of the token.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace is used to scope the query.

<a id="schema-token"></a>

### token property

Type: `"string"`. Computed, Sensitive.

Access token for the F5 Artifact Registry (FAR) This token can be used to authenticate with FAR when
pulling related images for Kubernetes bot infrastructure.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `expiration_time` | [expiration_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/properties/#schema-expiration_time) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/properties/#schema-namespace) |
| `token` | [token](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/properties/#schema-token) |

## Next pages

- [xcsh_artifact_registry_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/)
