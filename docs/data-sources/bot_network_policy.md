---
page_title: "xcsh_bot_network_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_network_policy landing."
---

# xcsh_bot_network_policy landing

<a id="canonical-0131231022333300-2321113231232321-2013223130013312-1311321020233021-2120000102000123-3300101133122113-2333131010121013-3020322100131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301002222201011-0231012121211202-0220102010003213-2222120321202101-3330113313201231-3323032310222323-0200312002030020-2023312333032013"></a>

## xcsh_bot_network_policy — xcsh_bot_network_policy / 223012022020 / 2

Breadcrumbs:

- xcsh_bot_network_policy

Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy.
configuration. (read-only data source)

<a id="canonical-0021102213301312-1133302131120120-1310131002213003-0112211021103113-1201132100022103-1210210320323330-1100323023331233-3320122121232311"></a>

## Prerequisites — xcsh_bot_network_policy / 223012022020 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3220131300323202-2012200311311220-2320331032103301-3123022100330223-3200321120013313-3032022331230011-1233013210031122-3331221012100120"></a>

## Minimal configuration — xcsh_bot_network_policy / 223012022020 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```

<a id="canonical-2310022030321313-3301012022012102-1222303010310213-0103122001212313-1233211301201020-2112113010122323-0012232131122013-2310112022021121"></a>

## Root configuration — xcsh_bot_network_policy / 223012022020 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1131003100003030-2203313211103203-1200210311323003-0212222112300230-2310113103131303-0112203020132120-0023203300102032-2121323320100330"></a>

## Next pages — xcsh_bot_network_policy / 223012022020 / 6

- [Property reference](../guides/data-sources--bot_network_policy--reference--group-001.md#canonical-1000220123222103-1311203111200122-1310233113223010-2303233320132230-2023011120122200-0020202331302112-2201330132003322-1123131221011230)
- [Examples](../guides/data-sources--bot_network_policy--examples--group-001.md#canonical-0002022322030223-0312030020301110-0013302131302112-1230333220320331-3031320232323312-3121201120303031-1030210030322301-1330332213001121)
