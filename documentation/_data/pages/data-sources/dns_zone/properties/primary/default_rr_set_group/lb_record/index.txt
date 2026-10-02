---
page_title: "primary.default_rr_set_group.lb_record"
subcategory: "DNS"
description: "DNS Load Balancer Record."
xcsh_docs: {"aliases": ["primary default rr set group lb record"], "body_bytes": 3096, "body_sha256": "sha256:1afe4c30a04e24ccde56182aec69854e79a262e582100b43fb7185b72a260663", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record:value"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "lb_record"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-primary--default_rr_set_group--lb_record--name", "description": "Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name and not a subdomain of a subdomain.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "lb_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record:value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["primary", "default_rr_set_group", "lb_record", "value"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DNS Load Balancer Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.lb_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- primary.default_rr_set_group.lb_record

<a id="section"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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

<a id="schema-primary--default_rr_set_group--lb_record--name"></a>

### name property

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group.lb_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
