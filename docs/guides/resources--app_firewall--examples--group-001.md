---
page_title: "xcsh_app_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall examples."
---

# xcsh_app_firewall examples

<a id="canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- Examples

<a id="canonical-1002333111330313-3202303331212002-1210121210022330-0212223200021031-2221302002032313-3011122031332110-0300123023032010-3011012212323100"></a>

### Complete configurations for `xcsh_app_firewall`

- [Ai enhancements](resources--app_firewall--examples--group-001.md#canonical-1123313312132321-2133231331301002-1323202312013033-0102332212012011-3001323210230101-0322221032200212-1220112133010330-2000021011020330): valid configuration.

- [All attributes](resources--app_firewall--examples--group-001.md#canonical-0011002223310300-3001203030222321-3132031200030121-0333331113110203-3133030032022200-1031321121003311-3333030110220121-1320112220331111): valid configuration.

- [Allowed response codes](resources--app_firewall--examples--group-001.md#canonical-0203213320003220-1301032123212000-1332232001222030-1103103320122132-3122323103311322-3303032233323020-0212323202330310-3020113212303031): valid configuration.

- [Blocking](resources--app_firewall--examples--group-001.md#canonical-0332302023310211-0300010022330303-3303322211000332-2011101230233310-0100123200232311-2303111230001210-2231313003213022-2203212330020023): valid configuration.

- [Bot protection](resources--app_firewall--examples--group-001.md#canonical-0211211133000211-3001323131330100-2101220020033303-3332001321322033-1322203210020101-3323033020022330-2003201020112222-3321210302020120): valid configuration.

- [Custom blocking page](resources--app_firewall--examples--group-001.md#canonical-3030200331103100-3101110011212013-0333230030101303-0201233212122322-0102010113030120-1031331022311321-3312212132122333-3303112223121131): valid configuration.

- [Detection settings](resources--app_firewall--examples--group-001.md#canonical-2021221133300300-1122033021031110-1100210202002123-3032001223103320-3110131122302213-3313032232122210-2333311101020021-3133113010310323): valid configuration.

- [Disable anonymization](resources--app_firewall--examples--group-001.md#canonical-2110210122122203-3323213231132120-1111132123303212-2213203320203301-3002311131123200-2022321322101310-2123122213001130-0231121221330132): valid configuration.

- [Monitoring](resources--app_firewall--examples--group-001.md#canonical-0133102032231101-0031312110101210-2203301323232223-0002123123202002-3200031131313201-2113121321311110-0201001330210132-1310013001330010): valid configuration.

- [Resource](resources--app_firewall--examples--group-001.md#canonical-2212130033200011-3131123033222113-1312132301000221-3321311011213200-3103222201202231-3000320132003002-2033331001201113-3331122230233023): valid configuration.

- [With labels](resources--app_firewall--examples--group-001.md#canonical-2030311032102122-1000300211002213-0320313220310320-3003132331232001-0300220121323123-1113011310213301-2020320301333203-0321113100013212): valid configuration.

- [With updated labels](resources--app_firewall--examples--group-001.md#canonical-3220033311201323-0220231123202100-3320212210212113-3220330013122010-1032333323302330-0330221322212032-2110300111313030-1023220322032102): valid configuration.

<a id="canonical-1123313312132321-2133231331301002-1323202312013033-0102332212012011-3001323210230101-0322221032200212-1220112133010330-2000021011020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Ai enhancements example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Ai enhancements

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/ai-enhancements.tf`; digest `sha256:8a5b5775deccd3da5f7c4e690d3bca03ede520f6e7ee1e00db0fa32e73bad912`.

```terraform
# AiEnhancements — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}

  enable_ai_enhancements {
    mitigate_high_risk_action = {}
  }
}
```

<a id="canonical-0011002223310300-3001203030222321-3132031200030121-0333331113110203-3133030032022200-1031321121003311-3333030110220121-1320112220331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/all-attributes.tf`; digest `sha256:15478fdf93d3310174c8b51758e8f4e5d947ab3506a55afff7d878110c0c353d`.

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

resource "xcsh_app_firewall" "test" {
  name        = "example"
  namespace   = "system"
  description = "Full attributes test"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
  }

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}
}
```

<a id="canonical-0203213320003220-1301032123212000-1332232001222030-1103103320122132-3122323103311322-3303032233323020-0212323202330310-3020113212303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Allowed response codes example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Allowed response codes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/allowed-response-codes.tf`; digest `sha256:9fac3f5328e4e5b2882046677e97a6cbe0fb6c753fef26f8f52956cc720b9fe4`.

```terraform
# AllowedResponseCodes — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}

  allowed_response_codes {
    response_code = [200, 204, 301, 302]
  }
}
```

<a id="canonical-0332302023310211-0300010022330303-3303322211000332-2011101230233310-0100123200232311-2303111230001210-2231313003213022-2203212330020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Blocking example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Blocking

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/blocking.tf`; digest `sha256:ef7e8d73d30409b783a8406919290ea0de139474cbcea62a1582fcfe0f55c136`.

```terraform
# Blocking — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  # Use default detection settings
  default_detection_settings = {}

  # Blocking mode - actively block malicious requests
  blocking = {}

  # allow_all_response_codes / use_default_blocking_page / default_bot_setting /
  # default_anonymization are server-default oneof markers the provider import-suppresses.
  # Declaring them makes the config import-unclean (config has them, imported state does
  # not), so they are intentionally omitted here — the server still materializes them.
}
```

<a id="canonical-0211211133000211-3001323131330100-2101220020033303-3332001321322033-1322203210020101-3323033020022330-2003201020112222-3321210302020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Bot protection example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Bot protection

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/bot-protection.tf`; digest `sha256:6e9c8727741ef612628e60f244a11489e11b01d40ff73ee72db9aa6ea5f61164`.

```terraform
# BotProtection — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_anonymization      = {}

  bot_protection_setting {
    good_bot_action       = "REPORT"
    malicious_bot_action  = "BLOCK"
    suspicious_bot_action = "REPORT"
  }
}
```

<a id="canonical-3030200331103100-3101110011212013-0333230030101303-0201233212122322-0102010113030120-1031331022311321-3312212132122333-3303112223121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Custom blocking page example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Custom blocking page

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/custom-blocking-page.tf`; digest `sha256:f2df3fea5ad8e1732d977424427aaf5de894118f0229d9185c0686db929c5462`.

```terraform
# CustomBlockingPage — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  default_bot_setting        = {}
  default_anonymization      = {}

  blocking_page {
    blocking_page = "https://example.com/blocked.html"
    response_code = "Forbidden"
  }
}
```

<a id="canonical-2021221133300300-1122033021031110-1100210202002123-3032001223103320-3110131122302213-3313032232122210-2333311101020021-3133113010310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Detection settings example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Detection settings

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/detection-settings.tf`; digest `sha256:50e5c6c8e8107cae5d8fe6c732ac501ab51251e38c9d54deab9761fd88fcb6f2`.

```terraform
# DetectionSettings — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  allow_all_response_codes  = {}
  blocking                  = {}
  use_default_blocking_page = {}
  default_bot_setting       = {}
  default_anonymization     = {}

  detection_settings {
    default_violation_settings = {}
    default_bot_setting        = {}
    enable_suppression         = {}
    enable_threat_campaigns    = {}
    signature_selection_setting {
      high_medium_accuracy_signatures = {}
      default_attack_type_settings    = {}
    }
  }
}
```

<a id="canonical-2110210122122203-3323213231132120-1111132123303212-2213203320203301-3002311131123200-2022321322101310-2123122213001130-0231121221330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Disable anonymization example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Disable anonymization

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/disable-anonymization.tf`; digest `sha256:4e7ac5ea3e4dfe4980ba8b5bf719912588c340253be3f414343af180a9e177fc`.

```terraform
# DisableAnonymization — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}

  disable_anonymization = {}
}
```

<a id="canonical-0133102032231101-0031312110101210-2203301323232223-0002123123202002-3200031131313201-2113121321311110-0201001330210132-1310013001330010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Monitoring example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Monitoring

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/monitoring.tf`; digest `sha256:c61bb69f05a187d2bdd987f0e88218a9e36bc46a2b425742c28302526d082474`.

```terraform
# Monitoring — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  default_detection_settings = {}
  allow_all_response_codes   = {}
  monitoring                 = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}
}
```

<a id="canonical-2212130033200011-3131123033222113-1312132301000221-3321311011213200-3103222201202231-3000320132003002-2033331001201113-3331122230233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/resource.tf`; digest `sha256:03e5e1d1ea7b4ce0a0005da7cf6bf27eeca5913c46ae5552f44fffd430a9cb75`.

```terraform
# AppFirewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppFirewall configuration
resource "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}
```

<a id="canonical-2030311032102122-1000300211002213-0320313220310320-3003132331232001-0300220121323123-1113011310213301-2020320301333203-0321113100013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/with-labels.tf`; digest `sha256:51f4bc6ac4257ec69ec14e7e0c37f3b9cd582e2461dff7632a889f943d97b42f`.

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

resource "xcsh_app_firewall" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test application firewall"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Use default detection settings
  default_detection_settings = {}

  # Allow all response codes
  allow_all_response_codes = {}

  # Blocking mode
  blocking = {}

  # Use default blocking page
  use_default_blocking_page = {}

  # Use default bot settings
  default_bot_setting = {}

  # Use default anonymization
  default_anonymization = {}
}
```

<a id="canonical-3220033311201323-0220231123202100-3320212210212113-3220330013122010-1032333323302330-0330221322212032-2110300111313030-1023220322032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With updated labels example

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Examples](resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- With updated labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/with-updated-labels.tf`; digest `sha256:0bbaf48e9cd535c6fd630956b0756feb6f88090c3cb140fc62c351e03346e5cd`.

```terraform
# WithUpdatedLabels — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test application firewall"

  labels = {
    environment = "staging"
    team        = "platform"
  }

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}
}
```
