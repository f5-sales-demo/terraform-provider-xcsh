---
page_title: "ingress_rules.applications"
subcategory: ""
description: "Application protocols like HTTP, SNMP."
xcsh_docs: {"aliases": ["ingress rules applications"], "body_bytes": 1834, "body_sha256": "sha256:86b750f8e80f8740ce0231d8bba1309508e3e3f505a12dac8a6ac075d7189c5c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "parent_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "path": "documentation/data-sources/network_policy_view/properties/ingress_rules/applications/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3332100121000323-0023323322202301-2001212311212200-3120203333313201-3313131212113023-3221133103211103-2132013313322301-3212330332120302", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_rules", "applications"], "schema_version": 1, "sections": [{"aliases": ["applications"], "anchor": "schema-ingress_rules--applications--applications", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "applications", "applications"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/applications/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Application protocols like HTTP, SNMP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.applications

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/)
- ingress_rules.applications

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="schema-ingress_rules--applications--applications"></a>

### applications property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

## Next pages

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
