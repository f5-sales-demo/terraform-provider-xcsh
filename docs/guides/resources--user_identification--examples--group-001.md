---
page_title: "xcsh_user_identification examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification examples."
---

# xcsh_user_identification examples

<a id="canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- Examples

<a id="canonical-3303100000130033-1311020103031021-1130302332302303-1331032211232002-0231223131112130-2102112303233212-0213201222322010-2221202033331032"></a>

### Complete configurations for `xcsh_user_identification`

- [All attributes](resources--user_identification--examples--group-001.md#canonical-3132212113333202-3100231220311001-0031211121103213-1131202230221233-3000022232110010-3111223321301120-1313113331211200-2111103302201101): valid configuration.

- [Cookie](resources--user_identification--examples--group-001.md#canonical-0213113321013213-3210122030022033-0311331313313200-1000233122011303-2013121221231230-0233032300231222-3222020122113321-2330231013231033): valid configuration.

- [Http header](resources--user_identification--examples--group-001.md#canonical-1103030012312232-2221013300210100-1212002232003112-2112332000233322-2222333313011310-1301010001033031-2133000100002031-3112031303020232): valid configuration.

- [Resource](resources--user_identification--examples--group-001.md#canonical-0021030033310023-3323121131320120-2331223202320121-1111003123223123-2213223213002321-0121003220303213-3100001300010001-2122000223301302): valid configuration.

- [Tls fingerprint](resources--user_identification--examples--group-001.md#canonical-0331220123100321-2232002120000133-0000022303231133-3302321233120223-2331211231102231-1232333102133030-3300300322001211-2131300102011010): valid configuration.

- [With annotations](resources--user_identification--examples--group-001.md#canonical-3311100230323221-0011100032132010-0321103010011311-2010100032303113-0111130122101232-3121030303201011-2223132312110233-0211103212022323): valid configuration.

- [With description](resources--user_identification--examples--group-001.md#canonical-3230320200122111-3032130232331310-2113130211013131-1033013122310231-3211313200033213-1100023221222131-3030102212211220-0013013203022123): valid configuration.

- [With labels](resources--user_identification--examples--group-001.md#canonical-3100123301012123-0302233020220013-2330333113221121-2220131322332022-0300103132113031-0302100332110313-0110012022133332-1221001131322231): valid configuration.

- [With rules](resources--user_identification--examples--group-001.md#canonical-0112131323123300-3322213203131032-2000022012120212-0032110123010012-1330323213221220-1033313100231122-3231222212123031-0331223331331022): valid configuration.

<a id="canonical-3132212113333202-3100231220311001-0031211121103213-1131202230221233-3000022232110010-3111223321301120-1313113331211200-2111103302201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/all-attributes.tf`; digest `sha256:678a876b16600386f26af7457406ace3f301d5ca150aba0447d4b245d2b5a900`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test user identification with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-0213113321013213-3210122030022033-0311331313313200-1000233122011303-2013121221231230-0233032300231222-3222020122113321-2330231013231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Cookie example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- Cookie

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/cookie.tf`; digest `sha256:5603d5f0bc4ec0397625fee3d3eb29ea6bfe17b42e387a451f62ff082893af26`.

```terraform
# Cookie — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    cookie_name = "session_id"
  }
}
```

<a id="canonical-1103030012312232-2221013300210100-1212002232003112-2112332000233322-2222333313011310-1301010001033031-2133000100002031-3112031303020232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http header example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- Http header

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/http-header.tf`; digest `sha256:b19736efd2e6603701aad68c2c3118fc31fc2264595ed7e6f49cbd6de0dc5c92`.

```terraform
# HttpHeader — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    http_header_name = "X-Forwarded-For"
  }
}
```

<a id="canonical-0021030033310023-3323121131320120-2331223202320121-1111003123223123-2213223213002321-0121003220303213-3100001300010001-2122000223301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/resource.tf`; digest `sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2`.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

<a id="canonical-0331220123100321-2232002120000133-0000022303231133-3302321233120223-2331211231102231-1232333102133030-3300300322001211-2131300102011010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Tls fingerprint example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- Tls fingerprint

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/tls-fingerprint.tf`; digest `sha256:18aace91bd93f5201f1f6ed23d33d48d6c8f66a90c3a44614cea4e6f6f5a4a09`.

```terraform
# TlsFingerprint — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    tls_fingerprint = {}
  }
}
```

<a id="canonical-3311100230323221-0011100032132010-0321103010011311-2010100032303113-0111130122101232-3121030303201011-2223132312110233-0211103212022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With annotations example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-annotations.tf`; digest `sha256:903200cccd3aa8c5abe2eebfb30b2e2efe55bae4a37636c3dae4de46ed18a25f`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  annotations = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-3230320200122111-3032130232331310-2113130211013131-1033013122310231-3211313200033213-1100023221222131-3030102212211220-0013013203022123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With description example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-description.tf`; digest `sha256:28e44c81af88136d36f13495510529b754168c2724c3ce7cec3eeb9297eaa127`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "example-value"

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-3100123301012123-0302233020220013-2330333113221121-2220131322332022-0300103132113031-0302100332110313-0110012022133332-1221001131322231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-labels.tf`; digest `sha256:ba484df9375d0da693ea56c4aaf4288ad5030a88885849721fac1dec608ba5a3`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-0112131323123300-3322213203131032-2000022012120212-0032110123010012-1330323213221220-1033313100231122-3231222212123031-0331223331331022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With rules example

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112)
- [Examples](resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- With rules

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-rules.tf`; digest `sha256:b71ec64258f8256d380d65c27c4c945068a8b584275196cad8f9d2b7fcbb430f`.

```terraform
# WithRules — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "User identification with identification rules"

  rules {
    client_ip = {}
  }
}
```
