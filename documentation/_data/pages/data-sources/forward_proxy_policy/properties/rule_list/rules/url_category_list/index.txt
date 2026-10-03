---
page_title: "rule_list.rules.url_category_list"
subcategory: "Security"
description: "List of URL categories."
xcsh_docs: {"aliases": ["rule list rules url category list"], "body_bytes": 5834, "body_sha256": "sha256:667fd7f22bffcc4994ad3eb68765a3516a4925a416bab77289070c4a7cf24bb3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:url_category_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/url_category_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0233202033222111-2122021223212100-2100002000120303-2133131211221023-2302211312302313-1031130223233300-0112000112200222-3200310020303232", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "url_category_list"], "schema_version": 1, "sections": [{"aliases": ["rule list rules url category list url categories"], "anchor": "schema-rule_list--rules--url_category_list--url_categories", "description": "List of URL categories to be selected.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:url_category_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "url_category_list", "url_categories"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/url_category_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of URL categories.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.url_category_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.url_category_list

<a id="section"></a>

Type: `"single"`. Computed.

URL Category List Type. List of URL categories.

Upstream description:

List of URL categories.

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

<a id="schema-rule_list--rules--url_category_list--url_categories"></a>

### url_categories property

Type: `["list", "string"]`. Computed.

\[Enum:
UNCATEGORIZED|REAL\_ESTATE|COMPUTER\_AND\_INTERNET\_SECURITY|FINANCIAL\_SERVICES|BUSINESS\_AND\_ECONOMY|COMPUTER\_AND\_INTERNET\_INFO|AUCTIONS|SHOPPING|CULT\_AND\_OCCULT|TRAVEL|ABUSED\_DRUGS|ADULT\_AND\_PORNOGRAPHY|HOME\_AND\_GARDEN|MILITARY|SOCIAL\_NETWORKING|DEAD\_SITES|INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS|TRAINING\_AND\_TOOLS|DATING|SEX\_EDUCATION|RELIGION|ENTERTAINMENT\_AND\_ARTS|PERSONAL\_SITES\_AND\_BLOGS|LEGAL|LOCAL\_INFORMATION|STREAMING\_MEDIA|JOB\_SEARCH|GAMBLING|TRANSLATION|REFERENCE\_AND\_RESEARCH|SHAREWARE\_AND\_FREEWARE|PEER\_TO\_PEER|MARIJUANA|HACKING|GAMES|PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY|WEAPONS|PAY\_TO\_SURF|HUNTING\_AND\_FISHING|SOCIETY|EDUCATIONAL\_INSTITUTIONS|ONLINE\_GREETING\_CARDS|SPORTS|SWIMSUITS\_AND\_INTIMATE\_APPAREL|QUESTIONABLE|KIDS|HATE\_AND\_RACISM|PERSONAL\_STORAGE|VIOLENCE|KEYLOGGERS\_AND\_MONITORING|SEARCH\_ENGINES|INTERNET\_PORTALS|WEB\_ADVERTISEMENTS|CHEATING|GROSS|WEB\_BASED\_EMAIL|MALWARE\_SITES|PHISHING\_AND\_OTHER\_FRAUDS|PROXY\_AVOIDANCE\_AND\_ANONYMIZERS|SPYWARE\_AND\_ADWARE|MUSIC|GOVERNMENT|NUDITY|NEWS\_AND\_MEDIA|ILLEGAL|CONTENT\_DELIVERY\_NETWORKS|INTERNET\_COMMUNICATIONS|BOT\_NETS|ABORTION|HEALTH\_AND\_MEDICINE|CONFIRMED\_SPAM\_SOURCES|SPAM\_URLS|UNCONFIRMED\_SPAM\_SOURCES|OPEN\_HTTP\_PROXIES|DYNAMICALLY\_GENERATED\_CONTENT|PARKED\_DOMAINS|ALCOHOL\_AND\_TOBACCO|PRIVATE\_IP\_ADDRESSES|IMAGE\_AND\_VIDEO\_SEARCH|FASHION\_AND\_BEAUTY|RECREATION\_AND\_HOBBIES|MOTOR\_VEHICLES|WEB\_HOSTING\]
URL Categories. List of URL categories to be selected. Possible values are \`UNCATEGORIZED\`,
\`REAL\_ESTATE\`, \`COMPUTER\_AND\_INTERNET\_SECURITY\`, \`FINANCIAL\_SERVICES\`,
\`BUSINESS\_AND\_ECONOMY\`, \`COMPUTER\_AND\_INTERNET\_INFO\`, \`AUCTIONS\`, \`SHOPPING\`,
\`CULT\_AND\_OCCULT\`, \`TRAVEL\`, \`ABUSED\_DRUGS\`, \`ADULT\_AND\_PORNOGRAPHY\`,
\`HOME\_AND\_GARDEN\`, \`MILITARY\`, \`SOCIAL\_NETWORKING\`, \`DEAD\_SITES\`,
\`INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS\`, \`TRAINING\_AND\_TOOLS\`, \`DATING\`, \`SEX\_EDUCATION\`,
\`RELIGION\`, \`ENTERTAINMENT\_AND\_ARTS\`, \`PERSONAL\_SITES\_AND\_BLOGS\`, \`LEGAL\`,
\`LOCAL\_INFORMATION\`, \`STREAMING\_MEDIA\`, \`JOB\_SEARCH\`, \`GAMBLING\`, \`TRANSLATION\`,
\`REFERENCE\_AND\_RESEARCH\`, \`SHAREWARE\_AND\_FREEWARE\`, \`PEER\_TO\_PEER\`, \`MARIJUANA\`,
\`HACKING\`, \`GAMES\`, \`PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY\`, \`WEAPONS\`, \`PAY\_TO\_SURF\`,
\`HUNTING\_AND\_FISHING\`, \`SOCIETY\`, \`EDUCATIONAL\_INSTITUTIONS\`, \`ONLINE\_GREETING\_CARDS\`,
\`SPORTS\`, \`SWIMSUITS\_AND\_INTIMATE\_APPAREL\`, \`QUESTIONABLE\`, \`KIDS\`,
\`HATE\_AND\_RACISM\`, \`PERSONAL\_STORAGE\`, \`VIOLENCE\`, \`KEYLOGGERS\_AND\_MONITORING\`,
\`SEARCH\_ENGINES\`, \`INTERNET\_PORTALS\`, \`WEB\_ADVERTISEMENTS\`, \`CHEATING\`, \`GROSS\`,
\`WEB\_BASED\_EMAIL\`, \`MALWARE\_SITES\`, \`PHISHING\_AND\_OTHER\_FRAUDS\`,
\`PROXY\_AVOIDANCE\_AND\_ANONYMIZERS\`, \`SPYWARE\_AND\_ADWARE\`, \`MUSIC\`, \`GOVERNMENT\`,
\`NUDITY\`, \`NEWS\_AND\_MEDIA\`, \`ILLEGAL\`, \`CONTENT\_DELIVERY\_NETWORKS\`,
\`INTERNET\_COMMUNICATIONS\`, \`BOT\_NETS\`, \`ABORTION\`, \`HEALTH\_AND\_MEDICINE\`,
\`CONFIRMED\_SPAM\_SOURCES\`, \`SPAM\_URLS\`, \`UNCONFIRMED\_SPAM\_SOURCES\`,
\`OPEN\_HTTP\_PROXIES\`, \`DYNAMICALLY\_GENERATED\_CONTENT\`, \`PARKED\_DOMAINS\`,
\`ALCOHOL\_AND\_TOBACCO\`, \`PRIVATE\_IP\_ADDRESSES\`, \`IMAGE\_AND\_VIDEO\_SEARCH\`,
\`FASHION\_AND\_BEAUTY\`, \`RECREATION\_AND\_HOBBIES\`, \`MOTOR\_VEHICLES\`, \`WEB\_HOSTING\`.
Defaults to \`UNCATEGORIZED\`.

Upstream description:

List of URL categories to be selected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
