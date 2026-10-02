---
page_title: "graphql_rules.metadata"
subcategory: "Load Balancing"
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["graphql rules metadata"], "body_bytes": 3528, "body_sha256": "sha256:3acb50ccc8285645baa4429056ad1b26a7547828e349e847a808471f42b8488f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:graphql_rules:metadata", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:graphql_rules", "path": "documentation/data-sources/http_loadbalancer/properties/graphql_rules/metadata/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "metadata"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-graphql_rules--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:graphql_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-graphql_rules--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:graphql_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/graphql_rules/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.metadata

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/graphql_rules/)
- graphql_rules.metadata

<a id="section"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

## Direct properties

<a id="schema-graphql_rules--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="schema-graphql_rules--metadata--name"></a>

### name property

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/graphql_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
