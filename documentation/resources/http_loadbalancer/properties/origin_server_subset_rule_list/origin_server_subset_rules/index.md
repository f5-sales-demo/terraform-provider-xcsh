---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules"
subcategory: "Load Balancing"
description: "Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct"
xcsh_docs: {"aliases": ["backend servers", "origin server subset rule list origin server subset rules", "origin servers", "upstream servers"], "body_bytes": 18792, "body_sha256": "sha256:14538b12a8421a5d1d4a2467b81292c5237e7d67eadc76981208700237f13a29", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:client_selector", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:none"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list", "path": "documentation/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:client_selector,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:client_selector,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:none", "type": "conflicts"}, {"anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--origin_server_subsets_action", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:RequiredListObjectAttributes:origin_server_subsets_action", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules"], "schema_version": 1, "sections": [{"aliases": ["any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--asn_list--as_numbers", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.asn_list:RequiredObjectAttributes:as_numbers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_list"], "syntax": "block", "type": "object"}, {"aliases": ["asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher:asn_sets", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:client_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--client_selector--expressions", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.client_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:client_selector", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "client_selector"], "syntax": "block", "type": "object"}, {"aliases": ["country codes"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--country_codes", "description": "List of Country Codes.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "country_codes"], "syntax": "attribute", "type": "list"}, {"aliases": ["ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--metadata--name", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:metadata", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin server subsets action", "origin servers", "upstream servers"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--origin_server_subsets_action", "description": "Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured in the origin pool are: 1. Add labels to origin servers 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "origin_server_subsets_action"], "syntax": "attribute", "type": "map"}, {"aliases": ["re name list"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--re_name_list", "description": "List of RE names for match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "re_name_list"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to..

Upstream description:

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("origin_server_subsets_action"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("client_selector",
    "none"),
  validators.ConflictingListObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
origin_server_subset_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/any_asn/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/any_ip/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/client_selector/): complete subsection reference.

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--country_codes"></a>

### country_codes property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_prefix_list/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/metadata/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/none/): complete subsection reference.

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--origin_server_subsets_action"></a>

### origin_server_subsets_action property

Type: `["map", "string"]`. Optional.

Add labels to select one or more origin servers.

Upstream description:

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--re_name_list"></a>

### re_name_list property

Type: `["list", "string"]`. Optional.

RE Names. List of RE names for match.

Upstream description:

List of RE names for match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/any_asn/)
- [origin_server_subset_rule_list.origin_server_subset_rules.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/any_ip/)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/)
- [origin_server_subset_rule_list.origin_server_subset_rules.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/client_selector/)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_prefix_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/metadata/)
- [origin_server_subset_rule_list.origin_server_subset_rules.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/none/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
