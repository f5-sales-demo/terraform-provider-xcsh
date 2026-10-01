---
page_title: "xcsh_workload_flavor landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor landing."
---

# xcsh_workload_flavor landing

<a id="canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210321203003331-3031120312030103-3020230311210202-2103110232132012-1132303121132131-2301313230120202-1121011220033120-1031102121113030"></a>

## xcsh_workload_flavor — xcsh_workload_flavor / 230020112303 / 2

Breadcrumbs:

- xcsh_workload_flavor

Manages workload\_flavor in F5 Distributed Cloud.

<a id="canonical-3311023011000313-2331303230130013-0312113131130030-3011223003313312-1231311221323001-3123212013200121-0023210323113222-0133312003001000"></a>

## Prerequisites — xcsh_workload_flavor / 230020112303 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2122122021302313-1331010321303103-3001001212131130-2010130001231132-3023231023202201-0013322133211101-2221210302333102-0133200003323013"></a>

## Minimal configuration — xcsh_workload_flavor / 230020112303 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WorkloadFlavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}
```

<a id="canonical-0020323003231012-3330102233031332-0120201333320233-3213302020130131-1231131000001231-1120112200132102-2121131313301232-0203013300201312"></a>

## Root configuration — xcsh_workload_flavor / 230020112303 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3320301323230311-0323110230122230-3123320002203131-3031300122232311-1001100202011321-3110033030023221-1331013101103110-0231203302133212"></a>

## Next pages — xcsh_workload_flavor / 230020112303 / 6

- [Property reference](../guides/resources--workload_flavor--reference--group-001.md#canonical-2201230231331130-0233333320310023-2033232302300302-2210032310122123-0221311000031233-1020311213100203-1113130201012033-0300212122012213)
- [Examples](../guides/resources--workload_flavor--examples--group-001.md#canonical-0021201222321131-3310011210200020-1122213013302121-3312000030010212-3122000022200221-2210211023201200-0113202133200222-2323301131323210)
- [Import](../guides/resources--workload_flavor--lifecycle--group-001.md#canonical-2323021122210002-2212110111100100-0233321233232203-1322003201220002-3323130031221320-1012200000333031-0303302222301133-2111222131130223)
- [Timeouts](../guides/resources--workload_flavor--lifecycle--group-001.md#canonical-3020121022001322-3032321133222012-0213120113330130-0120103313122201-3113203321010003-1202000113002230-3031321231220312-2333311002233132)
