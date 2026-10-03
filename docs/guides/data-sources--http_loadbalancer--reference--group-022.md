---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1202223110133102-3200303000300203-3031122132121121-3313131202030001-3211110202110320-1021031220220323-2300122030123332-3200001033023231"></a>

## name property — metadata / 321333101113 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3021112302112301-3332020021301202-0210202113201023-3111230032300110-3311012221133200-2130222023113320-3010210003011311-2312312312110233"></a>

## Next pages — metadata / 321333101113 / 6

- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302301330221333-1121010303022122-0231022211323031-0003100013311031-1103321113030101-1000200200131202-1300322111112210-3332032210021321"></a>

## policy_based_challenge.rule_list.rules.spec — spec / 101101011123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-0323031210333222-1233122213321110-2200220321203222-1332132121002211-3200103321213011-1011000132321300-0103131020221130-2213133001000221"></a>

Type: `"single"`. Computed.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-3200310313002133-0010311010212123-2312102101331200-0220220330102301-2110012303211210-2303022333112300-0031100002131010-2121220122333113"></a>

## Direct properties — spec / 101101011123 / 3

- [any_asn](data-sources--http_loadbalancer--reference--group-022.md#canonical-1110112001201001-2303203033000203-2130231020022321-0301000311232001-3303110003001103-3213333301100031-3233111003330311-2233200300222312): complete subsection reference.

- [any_client](data-sources--http_loadbalancer--reference--group-022.md#canonical-2131033212110203-1003133220231010-0120221011210230-2010103232223323-2201102302313333-1112211012012310-2021013231000133-1110323231322001): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-022.md#canonical-2221231100300023-2210020303223021-0031320102333202-2102200030032311-1121023002003232-1230320120100333-3030311210313130-0320312022112313): complete subsection reference.

- [arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-2102013312122211-0130022121103002-0113331230002130-1101210113331010-0303031332122311-1011021310002013-3013123220013223-3303110032133201): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221): complete subsection reference.

- [body_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-2130222302101313-2233302001002321-3322010333311213-1103213000211301-2133320123330332-2302210210112022-2321133102321132-2323201001222012): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-022.md#canonical-2000323223223323-2213103122032232-2032012302000332-1103023033311001-3303310000122131-2110300101110211-3021302032013102-2223201132322222): complete subsection reference.

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000): complete subsection reference.

- [disable_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-3020133121101200-2222013212131112-3111032310221132-2213303201101321-1300132123112201-3131333230313112-1220001302231030-2202110030031102): complete subsection reference.

- [domain_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3311122103211312-0002003200333203-1210023030112201-2031031301112330-3220213200303010-1323003002301022-1230010220321100-3120321133102132): complete subsection reference.

- [enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-2322120012012001-0300113331321331-0300333021300022-2012123332131133-3100110230212230-0001321233021031-1103301000033302-2331302011012102): complete subsection reference.

- [enable_javascript_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-0130320232310300-1003033303212310-1030202000201111-3022320032203012-1110200010232110-0320022322130201-0102110002122103-1202210113103203): complete subsection reference.

<a id="canonical-3033112312333232-0322231223101310-0020003010130031-2022222232313103-2022211111013020-0123221323323002-2203121000301023-0132100030302332"></a>

<a id="canonical-1313020010300122-2031132012230001-0030201302331320-3203230010133011-0033021323000132-1302011110013022-1113032033210013-1003112302031130"></a>

## expiration_timestamp property — spec / 101101011123 / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101): complete subsection reference.

- [http_method](data-sources--http_loadbalancer--reference--group-022.md#canonical-3020310301331122-1212000113011100-1023213212330331-3122103113210321-2002122203320320-2120121220311201-1333220030000032-3120130221023203): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-2102212001013000-1331112303230010-2332311101103003-1203312030331231-1110322311233210-0133212103001332-1212233031202211-1222020223230121): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-022.md#canonical-1130112213311222-2212120101303321-2312021201330112-1010000001222113-3331020212132323-1321320012001100-3330022120222112-2310110330301333): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-0022113021303313-2123000230233031-0231021111110113-2111313001202330-1220202120123310-1300312201020212-0013033111222322-0130112033123023): complete subsection reference.

<a id="canonical-1211132211210120-3213210023010200-3002321132111333-0331313211321323-2010133032310021-0000310331002030-3110003202120113-0011100123030230"></a>

## Next pages — spec / 101101011123 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](data-sources--http_loadbalancer--reference--group-022.md#canonical-1110112001201001-2303203033000203-2130231020022321-0301000311232001-3303110003001103-3213333301100031-3233111003330311-2233200300222312)
- [policy_based_challenge.rule_list.rules.spec.any_client](data-sources--http_loadbalancer--reference--group-022.md#canonical-2131033212110203-1003133220231010-0120221011210230-2010103232223323-2201102302313333-1112211012012310-2021013231000133-1110323231322001)
- [policy_based_challenge.rule_list.rules.spec.any_ip](data-sources--http_loadbalancer--reference--group-022.md#canonical-2221231100300023-2210020303223021-0031320102333202-2102200030032311-1121023002003232-1230320120100333-3030311210313130-0320312022112313)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- [policy_based_challenge.rule_list.rules.spec.asn_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-2102013312122211-0130022121103002-0113331230002130-1101210113331010-0303031332122311-1011021310002013-3013123220013223-3303110032133201)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-2130222302101313-2233302001002321-3322010333311213-1103213000211301-2133320123330332-2302210210112022-2321133102321132-2323201001222012)
- [policy_based_challenge.rule_list.rules.spec.client_selector](data-sources--http_loadbalancer--reference--group-022.md#canonical-2000323223223323-2213103122032232-2032012302000332-1103023033311001-3303310000122131-2110300101110211-3021302032013102-2223201132322222)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-3020133121101200-2222013212131112-3111032310221132-2213303201101321-1300132123112201-3131333230313112-1220001302231030-2202110030031102)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3311122103211312-0002003200333203-1210023030112201-2031031301112330-3220213200303010-1323003002301022-1230010220321100-3120321133102132)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-2322120012012001-0300113331321331-0300333021300022-2012123332131133-3100110230212230-0001321233021031-1103301000033302-2331302011012102)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-0130320232310300-1003033303212310-1030202000201111-3022320032203012-1110200010232110-0320022322130201-0102110002122103-1202210113103203)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- [policy_based_challenge.rule_list.rules.spec.http_method](data-sources--http_loadbalancer--reference--group-022.md#canonical-3020310301331122-1212000113011100-1023213212330331-3122103113210321-2002122203320320-2120121220311201-1333220030000032-3120130221023203)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-2102212001013000-1331112303230010-2332311101103003-1203312030331231-1110322311233210-0133212103001332-1212233031202211-1222020223230121)
- [policy_based_challenge.rule_list.rules.spec.path](data-sources--http_loadbalancer--reference--group-022.md#canonical-1130112213311222-2212120101303321-2312021201330112-1010000001222113-3331020212132323-1321320012001100-3330022120222112-2310110330301333)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-0022113021303313-2123000230233031-0231021111110113-2111313001202330-1220202120123310-1300312201020212-0013033111222322-0130112033123023)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1110112001201001-2303203033000203-2130231020022321-0301000311232001-3303110003001103-3213333301100031-3233111003330311-2233200300222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201233101110201-2222111333121210-3230122101201112-1132300013212323-3101213133313200-1110002011302202-3232301023320103-3110310223202223"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — any_asn / 011002110002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2220023100303032-0201020023121322-1020213323310013-3331211221230310-0230321033103003-2133322231303300-2002322302221203-3013030202323010"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0321321313210233-3133201323033001-1203320220023333-2311123200132022-1113232212230202-2210022002000211-0333200331301233-1112001021322202"></a>

## Direct properties — any_asn / 011002110002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310121331020300-1133331103312300-0112330203310122-2330033101221210-1002032031223000-1303122033010013-0033302012111210-2213330112323300"></a>

## Next pages — any_asn / 011002110002 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2131033212110203-1003133220231010-0120221011210230-2010103232223323-2201102302313333-1112211012012310-2021013231000133-1110323231322001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200202232131022-0201211211022202-1111300011303220-3200312003223201-0313033002112013-2331103200223123-2111101331032031-3120220212210001"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — any_client / 122333330333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-1320332133220011-1122301231023312-0212022112130303-3201123103131022-3032003100032233-3112231132220031-3302202301131321-2230113321023123"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1311000101223330-2030032330133222-3300032111110012-0223333112113320-2111322032303103-2133032001111103-2023230303133011-3323223331121321"></a>

## Direct properties — any_client / 122333330333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012001231112322-2131203013122321-1101000102321300-3330021220202310-1011001223213211-0010223101321303-1322121333132313-1312322030212300"></a>

## Next pages — any_client / 122333330333 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2221231100300023-2210020303223021-0031320102333202-2102200030032311-1121023002003232-1230320120100333-3030311210313130-0320312022112313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222001120113110-1011302203231102-3012023122021311-2231130102023201-3222000112021203-0302012312131311-3330223322310133-1122003323213012"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — any_ip / 031211230200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0020203310301203-0222332010332030-0333232231123232-2132212121232310-0302011223212231-1303200332210123-1112210021003002-0231110321210021"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2231020321021113-2222232011123130-0113110133013221-0103321312022102-2212201002313311-1211310220330203-2221132010333101-2012223200213122"></a>

## Direct properties — any_ip / 031211230200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120321132001232-2031121033221132-0113233133033033-1003323232310021-0303232230333011-3113223332333021-1323301233220021-1123021121200221"></a>

## Next pages — any_ip / 031211230200 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303200301301121-2121010210322130-1102202200223211-1220221231323032-2210221220320223-1112322113022111-3332233133030211-3111101301223222"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — arg_matchers / 312100110101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1303310122220111-1310201020311311-0022202211303133-0113322003103112-3000220033011012-3023310130301333-3330210133121100-1112101332002300"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3322320331333102-0312312201321220-0022312322010031-1020201302021203-1031222123021331-3330230301012322-1333323220102101-1002122022103000"></a>

## Direct properties — arg_matchers / 312100110101 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3133333330030322-2220000010213013-3110010100310230-3123211332302122-1120321320202223-0032130303003232-0133213021333212-3313220200332210): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0131230030203320-2020323201102121-0131233223113130-1220130231031001-2132221112102023-0303101232333333-1121320312132301-0332103121303333): complete subsection reference.

<a id="canonical-1300033201121020-3121120032303230-0300201201300211-2220110233022121-2302201331021310-3131220330122320-3101103312320132-3330122331231022"></a>

<a id="canonical-3020011133011323-2330103310112000-2022111201213000-1313332113300020-3113130010110101-2100033302010033-2202313102333011-1103110210100133"></a>

## invert_matcher property — arg_matchers / 312100110101 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-1103003032012030-2123231301121231-0303231223113312-2020201323002031-3102033021301203-0201210303120020-2231301110220133-2201020120000100): complete subsection reference.

<a id="canonical-2031110211220022-1302122011131010-0130201210101002-3300310203213300-0132223112030312-0110111231210230-0010112210232231-2122321321010303"></a>

<a id="canonical-3113033221323032-2333103332210132-2012023032231310-0010103132222203-0032001233132322-2303302102113113-0302331201111230-3112303301332221"></a>

## name property — arg_matchers / 312100110101 / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3321031101003223-3133030133201332-1200322330223210-2213012221220300-3201203331023302-3203022332313321-0220022013302310-0212230213102122"></a>

## Next pages — arg_matchers / 312100110101 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3133333330030322-2220000010213013-3110010100310230-3123211332302122-1120321320202223-0032130303003232-0133213021333212-3313220200332210)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0131230030203320-2020323201102121-0131233223113130-1220130231031001-2132221112102023-0303101232333333-1121320312132301-0332103121303333)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](data-sources--http_loadbalancer--reference--group-022.md#canonical-1103003032012030-2123231301121231-0303231223113312-2020201323002031-3102033021301203-0201210303120020-2231301110220133-2201020120000100)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3133333330030322-2220000010213013-3110010100310230-3123211332302122-1120321320202223-0032130303003232-0133213021333212-3313220200332210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120213003111321-1102133131121301-3121012320331300-2312120133323102-2123110011201111-0003031110013202-1331112130303232-1131133312023000"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — check_not_present / 312021020003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-3132110003131111-0312212112113120-2123100000223032-2112222012130030-3210212013002310-0220001332011111-0212120110300132-0122010221331310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1220300010222113-3220332033031230-1111330032301330-3330132232202003-0121231123222121-2012222121220303-2313210003131121-1332103020023103"></a>

## Direct properties — check_not_present / 312021020003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313032223303110-1230113200122202-3112303100033111-3233112220013210-3220331030111102-0002021123303000-1211121320103333-2113002302013102"></a>

## Next pages — check_not_present / 312021020003 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0131230030203320-2020323201102121-0131233223113130-1220130231031001-2132221112102023-0303101232333333-1121320312132301-0332103121303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013112020323211-1232032321102233-3122110021100000-2100231322313233-1021331310310323-3311230102231110-2201231200101000-1010322322033303"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — check_present / 013131031121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0331100002103000-2110300110112210-1123331221202330-2020030102120200-0212332011330000-0300311321000133-1020311210322232-2310123323211011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3111003033133303-2311103133132001-2003203230003310-0032133120010323-1321121002101033-3001002221000223-0223032330011222-3130103113121223"></a>

## Direct properties — check_present / 013131031121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200103230121002-3303331120001021-2021012002200213-3323010112331121-1000223320121010-1002022302110022-2330003113130331-3323310133320312"></a>

## Next pages — check_present / 013131031121 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1103003032012030-2123231301121231-0303231223113312-2020201323002031-3102033021301203-0201210303120020-2231301110220133-2201020120000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331210313213101-1231021323203320-3013130102133012-2321110313312110-1231303033103033-0020010021003111-3031232022100310-3111213130020213"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — item / 231021121001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-3022020131222231-2031302300231130-0211213323132330-1323110100220210-1331032231022312-0011001122211311-3001012022223202-0323212232003001"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3030022212033111-0120222230221013-0213312330303121-1010122313100130-1001130023102230-3102020331203001-0230020032102323-3323212113030330"></a>

## Direct properties — item / 231021121001 / 3

<a id="canonical-3301103100301020-3212311120100233-3100103331333012-2330322100220003-3131102233022132-1320123023013013-0203023131301122-3330202100022320"></a>

<a id="canonical-2231110020102221-0301211110323223-1322213012011031-1133012032132332-0303020331211012-3031201013232101-3221123130102203-1302231100011020"></a>

## exact_values property — item / 231021121001 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0203203023012221-0201011032122302-0021203220110220-2123021010103330-1233201011321320-1310210130311222-3122202023101100-1312202023102302"></a>

<a id="canonical-0233031021023222-1032003033112033-3102130203330221-0032113222202001-3310210233330221-3312230012300100-0221133031013031-2131323022211303"></a>

## regex_values property — item / 231021121001 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131021111021310-0332330311322210-2030032021111133-1113033310133001-3131020312010303-0221112132331010-3332032232312131-3022223311233122"></a>

<a id="canonical-2331131322121030-2321120020313200-3200023301131132-2121020233220230-3003303121101010-0322111110020300-1330030300331233-0000213132222023"></a>

## transformers property — item / 231021121001 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0033201220021313-0320301311313200-1001020102232233-3311220110200102-2102312132303123-0120003333121311-1022111010313010-0320310212230321"></a>

## Next pages — item / 231021121001 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2102013312122211-0130022121103002-0113331230002130-1101210113331010-0303031332122311-1011021310002013-3013123220013223-3303110032133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200011310113202-0012312100101001-2123320123133123-3310211030223310-3100202221023032-0122321333030322-3210331000121012-0230021211113001"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — asn_list / 331013010003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3232022000013312-3200233001020001-3320123322111222-1231113310203313-2101333021013102-2311020101201001-1112031030323333-1200013222133010"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-3311102303220323-3201003013300331-0220210121233010-0213101322202300-0111233102102022-1231112303212211-0021003133121002-0331330201200201"></a>

## Direct properties — asn_list / 331013010003 / 3

<a id="canonical-2133013201110012-3333233112010203-2020320002010323-2103112101333100-3220231003303111-3203032021010330-1132312223313103-1301122032301323"></a>

<a id="canonical-1233011123122223-1301112131332323-1210133132310012-1033112310203322-3221100023101303-1223133312223220-1220122010332023-0332311302300223"></a>

## as_numbers property — asn_list / 331013010003 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2001002100210102-0131313213203032-2000111301011002-3121300120111010-2323133301223311-1000311223233132-2113221110312100-2102220300012310"></a>

## Next pages — asn_list / 331013010003 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233333131133332-2033121022032332-0010031032132203-0202330313302010-0231003303230311-1112101130332210-0021102220203032-3032001300031133"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — asn_matcher / 223022202012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3220001330130201-0032103022002012-1123331232101010-1221001120032011-1020123223322103-1001023023201332-0032132302300111-3002211101233033"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-1101112323000233-3030102123221310-0203303321300213-1122213201001323-1302102100130003-3333111030023320-0213233120213310-2302011210023201"></a>

## Direct properties — asn_matcher / 223022202012 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-022.md#canonical-2330133203333230-3320312011011120-0003000002223301-2313202131110302-1230200012313310-1220131111110131-3210202332000010-0033330023301131): complete subsection reference.

<a id="canonical-0030110111001131-2130103030321320-0111012332033300-0110020131231310-3132303210032302-2031213231031013-2221233222000133-3003003230301322"></a>

## Next pages — asn_matcher / 223022202012 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-022.md#canonical-2330133203333230-3320312011011120-0003000002223301-2313202131110302-1230200012313310-1220131111110131-3210202332000010-0033330023301131)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2330133203333230-3320312011011120-0003000002223301-2313202131110302-1230200012313310-1220131111110131-3210202332000010-0033330023301131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111232103100230-2230330110202230-0033232121210313-3121031201003210-2012001033213101-3021233332032130-2103122202013212-3023002113233013"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — asn_sets / 132013320131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-0220113231231311-3230023011213210-0021020131323322-1222330131133220-2003100300122330-1321123131022010-1022210032032331-3223123101303302"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3132121221221232-3332120020302100-0312123322223303-1122102313331013-2012130320320003-2003301003322010-1213212323222022-3202320110131310"></a>

## Direct properties — asn_sets / 132013320131 / 3

<a id="canonical-3002322311301132-0302222011030120-1313222213123112-3033301030032031-0003302330202212-3202232333102110-0112301333032320-0202202321032221"></a>

<a id="canonical-2033022123213030-2133132232223203-3302303102032123-3300000223031323-3221013103230132-2100213211130103-3103021312303230-1003213021302220"></a>

## kind property — asn_sets / 132013320131 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020202301133111-3133123023032310-2013223133022230-1111122223211020-0212201230123211-2021033200012320-0113012132200031-0102013013233303"></a>

<a id="canonical-2033120231230213-3333302023003022-1033232132003012-2213323002130202-0133121023320223-1202320013210020-3103222113323333-1333311321330123"></a>

## name property — asn_sets / 132013320131 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200123230113132-0102113213020330-3300231123001022-0011201230133100-3101311002202213-3010220102322013-0233200310020020-1013321031330201"></a>

<a id="canonical-1200020010330132-2110331110322201-1020123032220030-1203002112333122-1201000102120222-2103200231332201-2010220122112220-1221033031322031"></a>

## namespace property — asn_sets / 132013320131 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  }
}
```

<a id="canonical-2133232101110300-0202330332101011-0011110203220021-3132013232313322-3220323303233233-3003013201122120-1300210321103202-3012331002101020"></a>

<a id="canonical-0221031001321112-0331311310212033-1030222203301220-0033232021313030-0033110012001131-1022233320233022-1222000212311313-0111021202201102"></a>

## tenant property — asn_sets / 132013320131 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3311301110223232-3312330003002330-1013303203123120-0230323233103210-0110212303010021-1300031323030222-0333223321120221-0011220132113313"></a>

<a id="canonical-1220331130310232-0202211233022103-2032313201201300-0023311122021022-2301333202230323-0322201113311111-2222020203212321-2113232332120033"></a>

## uid property — asn_sets / 132013320131 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2030310003031202-3303330330111021-2221130031103122-3020200232220301-2030123101032330-3112003032012130-0320312000331121-1123030310220001"></a>

## Next pages — asn_sets / 132013320131 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2130222302101313-2233302001002321-3322010333311213-1103213000211301-2133320123330332-2302210210112022-2321133102321132-2323201001222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312330013020012-0101133201031001-2201102303322012-3200121110322213-1303033321210111-2200300022101123-0223321033030233-0302030122123231"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — body_matcher / 110111023120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-2320203100130322-1320213032223100-2221301312313033-1311212031313012-2201313001110203-3111313110310312-2023302023023112-2332232001223133"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-1100221212132333-1102322010033030-2300111101003211-0102200003212002-2033223031012222-0203321113302010-1030320110012322-1220010022200001"></a>

## Direct properties — body_matcher / 110111023120 / 3

<a id="canonical-2203303212001303-0332132311332132-2203203200013223-2111300323110220-0300310333030032-2312020031011331-2132301113020221-1333120120311131"></a>

<a id="canonical-3102103333233101-1100022010121232-1321120000111301-0213132112311303-0120230103312113-1201220203133211-1312130231113302-2213001320102313"></a>

## exact_values property — body_matcher / 110111023120 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030032112313111-0303221323313010-1301110103030221-2313330003011220-2100301220000332-1112032032300133-2031031200233333-1121320313213301"></a>

<a id="canonical-1232102001032120-0211211031013202-3322003020113101-1010001131113100-3033002202233222-1311301212012230-1130021013322212-3103220302023010"></a>

## regex_values property — body_matcher / 110111023120 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101300010012111-2000033020112020-1120001113123222-1133030333101103-3331121321112312-1021321332102130-1112322113300131-0112102121002100"></a>

<a id="canonical-3020323331322013-2111232012033233-0223023233102020-3320213020010100-3212002023202221-3120123131113321-3333001200331302-2020001132202011"></a>

## transformers property — body_matcher / 110111023120 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020132120223120-0201230011322210-1032320021122211-1222121202001002-3321220012322221-1203112311203112-3232113120201303-1120231101302100"></a>

## Next pages — body_matcher / 110111023120 / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2000323223223323-2213103122032232-2032012302000332-1103023033311001-3303310000122131-2110300101110211-3021302032013102-2223201132322222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332322303120011-1223102300331313-0111230323101322-3112313011030333-1121123022132123-0131221212223023-2220302020121222-2112213311031131"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — client_selector / 033000313312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-2130302001110200-3113200111030130-3220302221122103-1222333022020330-1320201122303212-3120011121132110-3111011231213133-2313013013030333"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-3300121301010310-1200132321332210-3223102110303321-3113210332312012-3103101313233210-3031003010010231-1232233233122320-0030301230120003"></a>

## Direct properties — client_selector / 033000313312 / 3

<a id="canonical-2223022131233120-1023113210212333-0002310200010310-1030230112003130-2010000000220222-3213011021113011-3210031123231321-1302230020020233"></a>

<a id="canonical-3111023023112230-3212133213302310-2112101213030312-2021222211120232-0223310331332213-3231322212112101-3200322132033211-1231201233011031"></a>

## expressions property — client_selector / 033000313312 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2300131012312033-0330200211221130-2031330010331133-2033210120111313-3011112310032000-1231001122133330-1103031031101130-0200032210201132"></a>

## Next pages — client_selector / 033000313312 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020133021002310-3031222333111321-2030203201310132-3102220102020110-3210021012310320-3102132222132033-1210223203003231-1021203123131022"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — cookie_matchers / 002222321201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-0330220220022312-2113132102033122-2331130300201321-0300222122130130-0003031012231323-1300320231332032-3133302033001032-2013313121203232"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2332023333023232-2110322122130220-3232020131322321-2103002210303112-0020311012312303-1301331002100112-0110101233031321-1001202102220323"></a>

## Direct properties — cookie_matchers / 002222321201 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0031222222021030-1123230322221101-2312222211332320-2313313311300321-0111310300121311-0022223222230002-3211121110101033-1102303030211321): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0001102211232200-3033001330012222-2322132003313031-0323333322311333-0122022010012011-3300102132131101-1213222013232013-0032112232313211): complete subsection reference.

<a id="canonical-1311312002320120-0322333130000233-3231301101210100-0313023100330230-2310100223112133-3213232312110012-0201021211220222-0232332113103010"></a>

<a id="canonical-0013202321101121-0203102110210203-2132222211333200-1213302232333111-2110231010021002-2201210101001000-0313310211123201-2010110030102230"></a>

## invert_matcher property — cookie_matchers / 002222321201 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-3202023032200013-2112000112203020-3322321231230211-2221300323112312-0013131000131023-0033130310220021-1002032030203000-0000000320001032): complete subsection reference.

<a id="canonical-0012322202002230-1222120203333030-1331213032011032-3112332030013321-1320231323110212-0301322013311311-2321023123012333-1300231130313132"></a>

<a id="canonical-1301321321130102-0131030322010232-2102201111102233-3230022012211102-1021332002231001-0011010222011000-0301311212202121-0020222031323111"></a>

## name property — cookie_matchers / 002222321201 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3021311032002121-0001000033032010-2213201233130332-0220110011010332-1123111000233223-0013010320103122-0132223000002230-3232113120330121"></a>

## Next pages — cookie_matchers / 002222321201 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0031222222021030-1123230322221101-2312222211332320-2313313311300321-0111310300121311-0022223222230002-3211121110101033-1102303030211321)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0001102211232200-3033001330012222-2322132003313031-0323333322311333-0122022010012011-3300102132131101-1213222013232013-0032112232313211)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-022.md#canonical-3202023032200013-2112000112203020-3322321231230211-2221300323112312-0013131000131023-0033130310220021-1002032030203000-0000000320001032)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0031222222021030-1123230322221101-2312222211332320-2313313311300321-0111310300121311-0022223222230002-3211121110101033-1102303030211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202233111211111-3311021202102022-0001200110200201-0102121313131003-1301231011300122-3131223003211310-2331211311010210-1133130013130001"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — check_not_present / 201030103110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-0002322133132002-0221313101213013-2212122203310001-3213133121313322-0121022322220301-2133103110203131-3130123000210023-1022001302200130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0032112210121003-3212102322022001-1203100201230122-1132222031321333-2210113130311110-3011331030222122-1001310030203220-2302322211320100"></a>

## Direct properties — check_not_present / 201030103110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112203310233211-2203320101311030-3230131222020032-1311203031200122-3222220132030002-1103221330322200-0002223023130220-0131203231223222"></a>

## Next pages — check_not_present / 201030103110 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0001102211232200-3033001330012222-2322132003313031-0323333322311333-0122022010012011-3300102132131101-1213222013232013-0032112232313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100123312231131-0313103310122021-1001103003201033-0231100310320220-3020001230323212-3300021010231302-0232231021113200-3013332221022320"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — check_present / 201123312231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2010030212220022-2033103212331323-3231203213032133-0110202010331300-2330201311033121-2001030133011311-1222011310300030-0233022312322120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1123331013010130-0031220222030232-0121222221233022-3100010202211031-2102302131102201-1230113123221100-3310320323013032-1232230033033323"></a>

## Direct properties — check_present / 201123312231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033020103233311-3020323203131320-3001031012030022-3002311011132032-0111131011120101-3321333110133030-2210213233330312-2232003101001221"></a>

## Next pages — check_present / 201123312231 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3202023032200013-2112000112203020-3322321231230211-2221300323112312-0013131000131023-0033130310220021-1002032030203000-0000000320001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011231310301322-3022200102121321-1323100222111011-0032303123113100-2001302310303112-2300013302132110-1313230331231101-2010212320120000"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — item / 302303021122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-3233123310212100-2013012010023022-3112121231332231-0213331010110013-3133311230303120-3011121033211132-1031213120010201-1311132222230202"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3112121000223303-2001311013012032-2100231300330131-3200020103101312-1113320221010011-0101321111100200-1231000310313333-3030231102230013"></a>

## Direct properties — item / 302303021122 / 3

<a id="canonical-2023130201232133-2333122012301232-2223131333100213-2201130130023212-0002012013310200-1332303111022302-1112113332123131-2000310331331010"></a>

<a id="canonical-2213312311200032-3102121113330203-2122021002221111-0003230011100033-1122033133320331-1323003202333301-1132013131331120-2321213011302132"></a>

## exact_values property — item / 302303021122 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2102100023332302-1032201012310102-1102333020210133-0112030021322001-2323003332321331-0100323312331202-2202022301200020-3002302203011330"></a>

<a id="canonical-3213202300021023-2100300033311232-3003021031221133-2100201132012113-3331301230010110-1321123301002300-1303132112020312-1320210300111102"></a>

## regex_values property — item / 302303021122 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031332203333332-2133100332103003-1002212300200200-3002023021023033-2303301323101003-0212132031310210-3120123100321033-2222023113001310"></a>

<a id="canonical-3133203303211202-1020031123210211-2333333103020031-2330101302201213-2310333011102010-2233023021302130-3312203001211100-0010012333202011"></a>

## transformers property — item / 302303021122 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302332021330120-2033103310233211-0311212012010130-1201011232021030-3333223203332331-2230120022113023-2022001102100203-3311021230102110"></a>

## Next pages — item / 302303021122 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3020133121101200-2222013212131112-3111032310221132-2213303201101321-1300132123112201-3131333230313112-1220001302231030-2202110030031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211013322001202-1221322013303302-2230230122212123-3003031022132223-1300332213330011-2330101231123110-1121011321322201-2010101130102102"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — disable_challenge / 010123102031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-0110103111021102-0323322010203133-2132313210011330-2321020002111130-1032201130303311-1002302322210030-2323010220013020-0012031033232200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable challenge.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2001012113121133-0230220331010113-1000130213221330-0133321100321121-3322111133011310-2121222113201023-1311000021330303-1312323103322032"></a>

## Direct properties — disable_challenge / 010123102031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320011220310123-0133023001331310-1030333320313100-3022112120311320-2031313132322103-1120102310211031-3132100013303221-3211023331123000"></a>

## Next pages — disable_challenge / 010123102031 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3311122103211312-0002003200333203-1210023030112201-2031031301112330-3220213200303010-1323003002301022-1230010220321100-3120321133102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113330120023001-3302101002132233-1333010110333303-3111211203220113-0020023301122202-0001301220131123-0122133123020223-1301231012011000"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — domain_matcher / 133122032031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-0022003301132120-2012330310002231-2003200321021121-0201323002001010-3303201301200300-3332022133331000-0201032010212100-3032311131202203"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3213031223102201-1332020020230213-1202121030302113-3131003002120230-0213231233001020-3113300213111303-0033222231221231-3220213310301022"></a>

## Direct properties — domain_matcher / 133122032031 / 3

<a id="canonical-3323313332033222-1132112031222230-3013120021313123-1100231331031322-1332222121113013-2032330110013203-3211323000202112-2111110303212302"></a>

<a id="canonical-1330100031120313-2112123023131301-0230013300201303-3333230300211010-3213222021100130-3332331320202201-1210221113333203-3332113321120131"></a>

## exact_values property — domain_matcher / 133122032031 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022021002231003-3233011220312233-0232022301113232-3133120131133312-3332332010002120-3310000222111202-0333131011213302-3220010121300322"></a>

<a id="canonical-0001323332233323-1330013323302131-2013221233103110-3120201201033302-2033231102202020-2320002123233111-3132212120210203-0130312031201102"></a>

## regex_values property — domain_matcher / 133122032031 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3321130333022003-1130133132330112-2200102102301102-3222231120201023-2012211213102012-0110333030330313-1013122121003323-0223232223320323"></a>

## Next pages — domain_matcher / 133122032031 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2322120012012001-0300113331321331-0300333021300022-2012123332131133-3100110230212230-0001321233021031-1103301000033302-2331302011012102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023030103231123-3031310000212120-0122131120100231-3311131101232300-3113301023332021-2231203320032231-1120213201001233-2300022230300031"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — enable_captcha_challenge / 223330011202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-1200021313030121-0123022221030023-2231322200300223-1033113202131013-3111100201002213-3012110212103020-1313132130011032-1003003223003230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable captcha challenge.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0011222123310012-1322113001102031-0013310313130020-0113000202331301-2221101200012000-0331303010331101-0330112003223310-2313111310022330"></a>

## Direct properties — enable_captcha_challenge / 223330011202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201113022100023-0031023001312000-2310020002020211-0100000332213030-0110320300331232-3010300120131313-2313030020000113-3132121113020022"></a>

## Next pages — enable_captcha_challenge / 223330011202 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0130320232310300-1003033303212310-1030202000201111-3022320032203012-1110200010232110-0320022322130201-0102110002122103-1202210113103203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221332031002132-1132332000130022-1203100212100103-2333002331113212-3033010010122202-2133212110001022-1202032001310011-2001200131230211"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — enable_javascript_challenge / 231023101232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-0020100031330330-1332120301020103-2000332231021221-1133231322203020-3300103301201321-2320300121001303-2230113030101202-1212333110213021"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1132023210002323-0123013333201021-2312032000223312-1132232020001231-3300120201202111-1312132120132320-2022320013213102-2122233303123313"></a>

## Direct properties — enable_javascript_challenge / 231023101232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031103032320102-0300102202331330-3130332110333233-0133122301330120-3311100122131332-2020112022202330-2033030302220311-1132130201103112"></a>

## Next pages — enable_javascript_challenge / 231023101232 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302023203010330-1202031123121312-2201103210320021-0120321331322000-3022320200211300-3232011122112020-1220001031111232-2310100230113012"></a>

## policy_based_challenge.rule_list.rules.spec.headers — headers / 232011231131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-1311012110101103-2022102003212133-2331122223022030-3031333303023030-2000311212003233-0321221133123001-0122233030012310-2112130211211113"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1330222102010120-1012200331013002-0321020333211210-0200013213100020-0322312222031013-2200102123230223-2212002013320010-2012221110221333"></a>

## Direct properties — headers / 232011231131 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3323311011222320-2030013323322120-3010020022223023-1033012230302202-2123232010313312-2303321011002130-3033302032033200-1001311122032221): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-2210223333011301-2122103132330213-2100323122032322-0313113112102202-1213302023032001-2133100231213332-2202203102130320-1301120032202120): complete subsection reference.

<a id="canonical-2100333321121233-2223111320233021-3020201312131233-2300022010302030-2221330113223011-1030312013033133-3220102333311323-0301221101221020"></a>

<a id="canonical-0331301012232201-0301120103201331-3113201112033312-0303123013100100-3003233302113313-0312222010121130-3210232331131101-2320333211332222"></a>

## invert_matcher property — headers / 232011231131 / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-0012330233033310-1033320213101320-3310023313021031-0020110031211302-1210131020003323-2203002032013100-0023220123332332-3121103332101233): complete subsection reference.

<a id="canonical-0123233031211021-1303113320011121-0303331133310131-2313030023003301-3231001322013003-2111300323131031-0300033333202023-0030321330012023"></a>

<a id="canonical-0303311303131213-2120121032030221-0300113311212121-0132021010310111-1232201312100010-1200101112013032-2120200010221103-3202010110223312"></a>

## name property — headers / 232011231131 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0322320133300001-1013000321232300-3132211321330232-2213221100312232-1001232130221200-1333131113211223-3213313222200033-1112201122303130"></a>

## Next pages — headers / 232011231131 / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3323311011222320-2030013323322120-3010020022223023-1033012230302202-2123232010313312-2303321011002130-3033302032033200-1001311122032221)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-2210223333011301-2122103132330213-2100323122032322-0313113112102202-1213302023032001-2133100231213332-2202203102130320-1301120032202120)
- [policy_based_challenge.rule_list.rules.spec.headers.item](data-sources--http_loadbalancer--reference--group-022.md#canonical-0012330233033310-1033320213101320-3310023313021031-0020110031211302-1210131020003323-2203002032013100-0023220123332332-3121103332101233)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3323311011222320-2030013323322120-3010020022223023-1033012230302202-2123232010313312-2303321011002130-3033302032033200-1001311122032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110023013300300-1311333020311233-2223322203332120-1100330031310000-1332232222113001-1232031103032111-3000213311322010-0333011010331311"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — check_not_present / 320331111120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-2100321020233210-0111101310220031-0220302021231012-0330003213130013-3020231020111122-0211112323213222-0222221131002101-2201313111232231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3102232323120303-1221012200030310-2301022133320311-0223001113232330-3223011313020131-1213031033122130-3312332131222313-2213333022212012"></a>

## Direct properties — check_not_present / 320331111120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003030213233300-3120021002310122-0003312222003300-2311010200133300-1010302030113023-1112222300230333-2222020323303023-3012101023111113"></a>

## Next pages — check_not_present / 320331111120 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2210223333011301-2122103132330213-2100323122032322-0313113112102202-1213302023032001-2133100231213332-2202203102130320-1301120032202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001203313003313-3203312131010022-1220301330011320-2323202220213131-1202011122132130-2130201311132230-0020203301300013-1203131301132100"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — check_present / 310123011023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-3120112032333021-2103301331013303-0032133231220022-0133103202232031-1313301031323221-3011112212301211-2020221333313310-3132131013101322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1131110213303222-3330212001100133-2210230203303202-0102030310332222-1211032131013002-1230002233130033-2232211223211123-3123031123112020"></a>

## Direct properties — check_present / 310123011023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001003203001303-3203023322101003-1020023233022102-2200221023003333-0311032113111211-1121032201011032-1012313010010122-0330112110220020"></a>

## Next pages — check_present / 310123011023 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0012330233033310-1033320213101320-3310023313021031-0020110031211302-1210131020003323-2203002032013100-0023220123332332-3121103332101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200013033230033-1212332203213332-2021003302123202-1212233202212213-0231033130330100-2102322013202000-3103230320100233-1300033001330000"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — item / 322310330221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-2311111320212101-3103013302120022-3211121113212023-2311102222330121-2212002030102023-3000233012112112-2103032011022110-1122332202020121"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3303321021311030-0303033231012302-0303212233011321-2201120213132321-1012110323211131-0011331102310010-0122203121110111-3012011013201030"></a>

## Direct properties — item / 322310330221 / 3

<a id="canonical-2312300113013031-2021333123103211-0003011013102210-2210322011000002-1332023210203220-2023211302221333-3223322033122000-1023032311202222"></a>

<a id="canonical-1232013113133201-2233032022203121-1033223032332332-1201230232130121-0300011322232231-0020130230311020-3001223103112200-3212330300000102"></a>

## exact_values property — item / 322310330221 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0210032212131333-1302231330220321-0323213001101113-0331000200001012-0203131102020033-0202103123203203-1200222211100333-0323103103201002"></a>

<a id="canonical-1112320311200301-1130313323000233-0020102120010002-0133011302112122-2230301330013002-3313202310130003-1023122210133221-2000223032100133"></a>

## regex_values property — item / 322310330221 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033021301212310-3133200203222313-0030310101113123-3020012311102211-2331033030031323-0131322313000102-2002300130132333-1321021333213121"></a>

<a id="canonical-2023310212000231-1213232211103111-3320303120013103-2011232122132113-0313321200320233-2221312123232210-3121032101002221-2022211100020333"></a>

## transformers property — item / 322310330221 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2100301322100002-2032203022033322-0000032303020113-3201221131313013-3133223030302130-2230212011122031-0012323203101222-0323131020202023"></a>

## Next pages — item / 322310330221 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3020310301331122-1212000113011100-1023213212330331-3122103113210321-2002122203320320-2120121220311201-1333220030000032-3120130221023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230200031013233-3232122210303320-0113112233111300-2223220113332030-1022100320000231-0133322322320330-1010021132323303-2011210031020003"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — http_method / 223020122023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-2232002321022103-2312101333311212-1330302022001232-3032110202111131-3001113301031203-0202212102210301-1230232231221000-3123110313301303"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-2302311133301000-3332021132102003-0322230202003303-0130201321120233-0003132330313003-0230121103001030-0011113330120320-3301101122101212"></a>

## Direct properties — http_method / 223020122023 / 3

<a id="canonical-0321020010113201-1320101301030102-1332310122010213-1223223013110221-3313330202011213-3230020111130301-1203020122200232-0203321013303320"></a>

<a id="canonical-2120202030200110-3300221213020220-3300033203013133-3201203100322021-2311202212223122-0023301100313122-2033032030210113-0032300321012222"></a>

## invert_matcher property — http_method / 223020122023 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-2313113103303130-2011212201302310-2102301300200023-0332221031032021-1330210033112213-1123300313311021-1312022103032313-2210212013302012"></a>

<a id="canonical-1020131010000031-2331132120323223-2310211030233123-0003103233130001-3211130223221132-2133233202100021-0231123122201223-2111030331201301"></a>

## methods property — http_method / 223020122023 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3012031000213201-1310322123333001-0032012200110200-2200212030022312-2110203303333223-3321030113123202-3312030131203231-3010132321111232"></a>

## Next pages — http_method / 223020122023 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022323020012211-2122333311313001-3200021300322213-2022321013032113-0102333311312013-1201231123022030-0222031133222031-0020000100312230"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — ip_matcher / 003200331222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-1100121103030130-0022012303133233-1002302123220211-3023000232011010-1333012003102303-1022102312100313-3121233300012311-2103123311313012"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-2211311113232321-1302033223111312-1010122103310213-1021302230020333-2230320232303112-3230112310023020-2120200220103202-3002003101332120"></a>

## Direct properties — ip_matcher / 003200331222 / 3

<a id="canonical-3210203211033230-2030003223202013-3013113013123012-3012011321130211-1100301312223031-1210032032313133-0331303202133302-2121300210332313"></a>

<a id="canonical-0331232231122302-3231213313102123-1122303100330121-2133301032103310-1330232133331321-0130011222323233-3202200333231320-3231101011200133"></a>

## invert_matcher property — ip_matcher / 003200331222 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-022.md#canonical-3321330223011200-2113221111001023-1010132303333213-0213102211233022-0213123033123312-0102321111200103-3310223201232203-2011311300020120): complete subsection reference.

<a id="canonical-1101011331210200-1221031120220000-0001021300023233-2121213211233110-0013223213112000-3302113131120203-1210100333113121-3322301031302103"></a>

## Next pages — ip_matcher / 003200331222 / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-022.md#canonical-3321330223011200-2113221111001023-1010132303333213-0213102211233022-0213123033123312-0102321111200103-3310223201232203-2011311300020120)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3321330223011200-2113221111001023-1010132303333213-0213102211233022-0213123033123312-0102321111200103-3310223201232203-2011311300020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120323211022302-2001220203312232-3222031303121100-1312320300313300-0202131231310123-3003031233110000-3122011223102320-0122331213333203"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — prefix_sets / 310103320220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-3323203121223011-0131011032032203-0122000333013321-2000223133312001-1032033231020330-2223033203323330-1210130032022322-2231130110322100"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1321320000121002-3023322313310031-1120120321111010-2303313223022333-2230331203231022-0330312031323131-1230012110012231-1323120310023332"></a>

## Direct properties — prefix_sets / 310103320220 / 3

<a id="canonical-3003213221022123-0132320132231202-0121312003011011-2312021200100323-0212113112213023-2013133321032210-0221310211110000-2003023113001230"></a>

<a id="canonical-2011211201001130-0121012030123003-1230200000330320-2313222102110311-0132232112210231-3003320120201302-2013221223020010-2321121203332133"></a>

## kind property — prefix_sets / 310103320220 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312331133201211-3322303113210320-0023111031132331-2123110312020230-2223220120322212-1130303103030020-0330120321003002-3121002100301002"></a>

<a id="canonical-3130031223211331-0212322133133211-2233313112023113-3311130133131323-1231300310230323-3311003010020233-2011231033020131-1003131022021001"></a>

## name property — prefix_sets / 310103320220 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0222023311032201-2111013111303203-0001322000203112-1001203030220123-3013222130111301-1220203002300230-2010023000103131-2310101320112202"></a>

<a id="canonical-2130012232321031-2100002033021113-2323303201231311-1331003020320322-0102301100310103-3033210012203120-1012012022002110-1303013103013112"></a>

## namespace property — prefix_sets / 310103320220 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  }
}
```

<a id="canonical-0030312002230022-2003302300112323-2220313220301121-1002120030003300-3333111000301310-2303233320132220-3302320213102221-2322223313212121"></a>

<a id="canonical-1221233033333010-2021033000303303-2102210333130022-3122001201100212-1010321003331100-3331102302112322-3201033330211323-3313303233023121"></a>

## tenant property — prefix_sets / 310103320220 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021033211331230-2301103233302111-3221112321332320-1333232330112120-2303132100202301-3301311323331300-1222110032003101-2231031222003113"></a>

<a id="canonical-1211122232001333-0122103213202332-3322330221022300-2101120003021312-0112021113112023-0211321021001311-1131331113032010-0221111313002231"></a>

## uid property — prefix_sets / 310103320220 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210023002330133-2211021200231222-2211320211300001-2133210130010232-0000032331200131-0013320212301302-0023220000303221-2221331322003012"></a>

## Next pages — prefix_sets / 310103320220 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2102212001013000-1331112303230010-2332311101103003-1203312030331231-1110322311233210-0133212103001332-1212233031202211-1222020223230121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223230221100311-3220003201212231-0201001220031302-0123200133002211-0102131231120313-1122120001230012-3121023132220312-1033233311310111"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — ip_prefix_list / 123311112131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-2100320012213103-1022210223101301-3120202322033231-0022120223112323-1011212000002002-1131001002312002-1303211202323123-1123312212201102"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-0212302232130231-3321303200222113-2230303211202311-0321330302131020-3002020110133121-2120012111110231-2223213113232122-3201321202100230"></a>

## Direct properties — ip_prefix_list / 123311112131 / 3

<a id="canonical-2003321100321130-3120202030220211-1232100012030231-3132201323300032-3230010110033300-0211233130020102-2103003133032103-1210132213213310"></a>

<a id="canonical-0300231201223130-0031330333222302-2212032022310102-3233033122320203-1120232132100001-0201033012200213-3210100331211113-3232022102313032"></a>

## invert_match property — ip_prefix_list / 123311112131 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-3001233132022020-1302112203132130-2003001321032003-3011103322011300-3132013131001021-3230023303101021-1300003223030303-2301331203001200"></a>

<a id="canonical-3220023200223002-1100201031103003-3130230011121010-0020003021320131-1333221221223313-0012231223133223-0131313010023033-2222011330231113"></a>

## ip_prefixes property — ip_prefix_list / 123311112131 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0011033023231103-1223101323013030-2013332313122132-1011220100223112-0331221222011003-3123300013021323-0302002312302202-3203322323310232"></a>

## Next pages — ip_prefix_list / 123311112131 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1130112213311222-2212120101303321-2312021201330112-1010000001222113-3331020212132323-1321320012001100-3330022120222112-2310110330301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120222330130302-0010321023000100-2113023110211011-1300103003310233-0331223111020332-1133000103202120-3022333002133311-1132333331313230"></a>

## policy_based_challenge.rule_list.rules.spec.path — path / 003101313020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1110133321100212-2313321020102220-1023231221000012-1101223310211010-3022103100202332-0332132010113121-1132231223003202-1002323012302102"></a>

Type: `"single"`. Computed.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-3002332220330020-2303000131110111-1200130033223000-0110212010002230-0013303312301100-3311133301000310-0023311021011010-2031332030213123"></a>

## Direct properties — path / 003101313020 / 3

<a id="canonical-2311201301010031-3131230123301300-0130002300010122-1311100331122011-3110023203001301-2002012310000303-0002210113312123-2300113232230322"></a>

<a id="canonical-1332022222123322-0130230233320200-0111100213321332-0131320200323220-1102332133223320-2212330230130223-1210321211133033-1301113311112121"></a>

## encoded_path_matcher property — path / 003101313020 / 4

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-1030033001131333-2120023013302212-0223232230003110-2100021131100101-3130030201033001-2220120211212212-3321321223302303-3303120310000200"></a>

<a id="canonical-3122220103231000-3011032310211101-0333030022210013-2023021230212330-3110130130322112-2020331320011033-3130331102021321-3211202313023113"></a>

## exact_values property — path / 003101313020 / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022232323003023-3120200330122032-0122021201101021-0312110232202312-3102311211203113-2002330132121113-2331001021223232-1132310223010010"></a>

<a id="canonical-0112001023203013-1111332211131332-2123123332020333-3020031021333030-3033300333032103-1110001202133312-2003210120310200-3303313030311021"></a>

## invert_matcher property — path / 003101313020 / 6

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-0122102103133303-0122003133113201-3112203033033100-3022312130032212-1210232302200202-3103210022313022-1203121022102112-3230221012112010"></a>

<a id="canonical-0211313110300101-2131322233203102-3320112200320101-2213023020310123-1232132000122022-0310323022001302-1113002112323213-1223013301321330"></a>

## prefix_values property — path / 003101313020 / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303010311222332-2312333120111100-0213012223000212-1121011210031010-2111333300301321-1320120100023232-3202320321203013-2212112021000200"></a>

<a id="canonical-3131223032211033-0100022112032203-2113213020212221-2303023201100222-0120211211321012-0222122122301210-1201002103003213-2323123113213321"></a>

## regex_values property — path / 003101313020 / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322033301131212-0100121220122233-2020211310220300-1323030330100222-2000112112212220-3223133232320200-3021001012231202-2132302332231001"></a>

<a id="canonical-0001221330201322-3222312003121031-2021113023031323-1210133202000000-0001200213222303-3000003330202131-0012211030201121-2313103322302312"></a>

## suffix_values property — path / 003101313020 / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111213330330132-0221322121333031-3132023330030202-3033032121132322-2331131222231210-2111320102001013-2030020311333233-3201230221231102"></a>

<a id="canonical-2010133212321123-1323301221220133-1320312031201012-0033333001000313-0102001301113312-0300222022131210-0122030011133332-2200022101103333"></a>

## transformers property — path / 003101313020 / 10

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101102231110122-3013232320012122-0312230311212123-0023220230011111-0113132011200001-1312211220021330-1010001213210010-1201130300330011"></a>

## Next pages — path / 003101313020 / 11

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220033323233112-0333301322301133-3222112003013001-0112131033012322-1300213032002222-0100122311201212-1001000213230203-1333321130311013"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — query_params / 013323313121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-1212322002201322-2110330103203321-2133302003022303-2102021332002111-0210003110120023-3223120321130013-3013111110202233-2013100003133003"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2111033121221221-3123232020232302-1330212110130202-3301023312302103-2213023230010213-2203010221312332-1102000233300131-3001313231331033"></a>

## Direct properties — query_params / 013323313121 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3301330120000030-1310320133003031-1323322323331212-2013121021013122-0012103223220123-1230300003331333-0131331320110330-2000320031133221): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-2101113120020302-3300300103223201-3002302132220303-3202211103002112-3100000203200320-3310123021313333-0300232230301000-3220231220001332): complete subsection reference.

<a id="canonical-2300102022001000-2131310232002203-2120203000222023-2112011010002331-0213111330230010-1022010301301212-3212202333032321-1331101231222123"></a>

<a id="canonical-0201211323322112-0233032030103111-0221003320202303-2003103110233030-0020001320223112-0111103203321211-3030030102201120-3020232010222113"></a>

## invert_matcher property — query_params / 013323313121 / 4

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023): complete subsection reference.

<a id="canonical-3202331131010032-1230011022331331-1011013021012101-3300102011303212-0233123010032213-3102221020310022-2313000331003321-1103300020010022"></a>

<a id="canonical-0222210130222000-1032023213000203-1302022310303023-1133100232201012-3102222201221233-0221230302132232-1220100001213223-3210312312020330"></a>

## key property — query_params / 013323313121 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3123320221002323-2321210033211301-3102122311322031-1133223231310113-1130203323032212-3001303212220221-0322232333231303-2110233311220130"></a>

## Next pages — query_params / 013323313121 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3301330120000030-1310320133003031-1323322323331212-2013121021013122-0012103223220123-1230300003331333-0131331320110330-2000320031133221)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-2101113120020302-3300300103223201-3002302132220303-3202211103002112-3100000203200320-3310123021313333-0300232230301000-3220231220001332)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](data-sources--http_loadbalancer--reference--group-022.md#canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3301330120000030-1310320133003031-1323322323331212-2013121021013122-0012103223220123-1230300003331333-0131331320110330-2000320031133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110233120301123-2121313333230023-2001302303303022-1320132112133211-1002230022333102-2232110230312122-2201210223310030-1020220313222010"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — check_not_present / 023013231203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-0232230022001023-3131300232303200-0223002003012131-1032102022110312-2223302212110001-2202232333223012-3021023212303201-0312020221233003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1111031033301310-3023321313103033-3333220133332130-1211103200011012-0000102101120310-0232122212030021-1110300330312130-2032213200333010"></a>

## Direct properties — check_not_present / 023013231203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120000121032001-3113000221311130-1011303110133030-1333211003022020-0130322013331010-0132220020022002-3012312213013133-2203020301222212"></a>

## Next pages — check_not_present / 023013231203 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2101113120020302-3300300103223201-3002302132220303-3202211103002112-3100000203200320-3310123021313333-0300232230301000-3220231220001332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321111203013002-0002131202012221-1132103211102020-1101001121203302-3211103023313011-3001202023012201-3202222031122121-3023323033302132"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — check_present / 300023121202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-0323233300032223-0113012322221010-2120302333111233-2133020220031223-0010112002201203-2123103310222022-2323300113101023-2100303131233323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3031002231301313-2232223213213021-3313131122113202-1023311122121331-0200231010131131-3122333201112210-2001013000212331-1111323120033310"></a>

## Direct properties — check_present / 300023121202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020001312212032-2322130230132323-0111221301321122-3221300131223323-3012120203233323-2033222023022022-3230122300212121-3101000231003301"></a>

## Next pages — check_present / 300023121202 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021330022103111-1122010313212212-2133230031211220-2201201212233312-0320213230311131-3223123112102132-2220323210213113-0110032122201201"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — item / 310121103202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-2232020031231011-0202320232223332-0020210213111212-2302200221132330-0320331013120122-3103020302203213-2331200013221112-1200112333211111"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-1022111202133200-0001000033203310-3333012220033020-2220221012313122-0130223022112333-0322022322011012-0132013221112003-0231333330122013"></a>

## Direct properties — item / 310121103202 / 3

<a id="canonical-1002301300202010-3221001102302100-0311123333020311-3133131201320131-2200121323223002-3133133210133311-1122011131020023-3311230003200203"></a>

<a id="canonical-0311121302203121-0202230322331123-2320323221001310-2210232122023313-1320311203031022-1103203000112222-0210303220222313-1231302023203010"></a>

## exact_values property — item / 310121103202 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3011021323113331-2221120130013331-1100213003030111-0131301000023322-0023032031302220-1112121030000213-1122221010331110-1120013120133202"></a>

<a id="canonical-1132201121222111-2202123122200221-1000103120311131-0112201203300330-3010203300123233-2011033133330331-1121013210121022-1012130113320230"></a>

## regex_values property — item / 310121103202 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123331013023120-3312210220322131-2310221312012301-1031202110023010-2131122233120332-2132302203031122-0131021130033021-2101231203030233"></a>

<a id="canonical-0231001331013120-1310300210130202-3301330322130313-0321313102302301-0311233003122002-3322023221020022-3031202023132012-2022210110031100"></a>

## transformers property — item / 310121103202 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2032112010102321-1303102033230330-0220231231102132-2121012011301213-1013300302210112-0212020203301200-1201213021320020-3311033102223033"></a>

## Next pages — item / 310121103202 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0022113021303313-2123000230233031-0231021111110113-2111313001202330-1220202120123310-1300312201020212-0013033111222322-0130112033123023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200201133221020-2311232130100333-3321323312120322-3033212100202100-2103000330300313-0200101033000121-1311001330131301-3201203101332131"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 122223212033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2032323200230031-2230223130230021-2010212122120201-0002213122002302-0010312312210032-3202122230132331-2201010133221113-3332233311003001"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-0233111113222030-0200231212201222-0123233113320032-1202202300302212-2211223102313010-1011103223012300-0022123121120131-3320223203220130"></a>

## Direct properties — tls_fingerprint_matcher / 122223212033 / 3

<a id="canonical-3023111302321211-0313202303300311-2302112222000332-0013130211323210-2232032321201120-1013023223132013-0311232302030303-1301333031212301"></a>

<a id="canonical-3023332102313121-3211020112102212-3021133222110100-0203202111301101-2111201222211221-3120203003312313-3333231130012032-2031031212310303"></a>

## classes property — tls_fingerprint_matcher / 122223212033 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212132003131222-3313321310030230-1302000012023213-0033010122123022-3230322303302130-3122312102023131-3321133000232133-1102213011123103"></a>

<a id="canonical-0110122230120022-2222322103300002-1223333210023302-2302311120112311-1011303322321211-2323110130031100-1231312110330212-2233223023321011"></a>

## exact_values property — tls_fingerprint_matcher / 122223212033 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3231121311030333-3213202031113011-0120021212113130-2011333232221211-1202132300020211-0013030023131230-1320111332221203-1011131233300303"></a>

<a id="canonical-3001030010213203-0222212313100321-2200212231123030-0213330232121310-1133233320033221-1222013211320121-0331000312323320-2320000322122031"></a>

## excluded_values property — tls_fingerprint_matcher / 122223212033 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3021102103010112-2233323122022120-1133330131323230-0212302331230312-1301011003100011-0332133320123120-1320211102303001-0031222210310202"></a>

## Next pages — tls_fingerprint_matcher / 122223212033 / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3113023112302332-0333303302133233-0010212202113330-1313232030220213-3320210011012203-0101333020233000-0132210002222132-1003010103121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003023112221111-3031211312012130-2121321111010111-1020230031322323-3113301010021113-3011032221332013-0221132302103203-2312013022123003"></a>

## policy_based_challenge.temporary_user_blocking — temporary_user_blocking / 212333313000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-0123013121122200-3132220222010300-2103133032302133-3121000310102231-3002010121022012-2311132030130002-0303020232223021-3130203130013210"></a>

Type: `"single"`. Computed.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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

<a id="canonical-2112121103121102-3131031203202222-0132322121002232-3210223203310222-3030000023211312-0003200032123131-3231103322033110-3320321110201212"></a>

## Direct properties — temporary_user_blocking / 212333313000 / 3

<a id="canonical-2002310133223303-2213332220030321-2030101103210323-0111013001321300-1320231023222223-3110211222023212-3330320101203001-2202110201130101"></a>

<a id="canonical-1002022023001120-3201100230333301-1233032021230023-3212020030031130-0311310003213333-2211133221302032-2102213300012201-1101233122131200"></a>

## custom_page property — temporary_user_blocking / 212333313000 / 4

Type: `"string"`. Computed.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1101111200033222-0131321220112001-2331010021131232-1101231312321300-2200321122113200-1302210331103230-1333100203220033-3321203020321123"></a>

## Next pages — temporary_user_blocking / 212333313000 / 5

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030232211130121-3313011233323231-3301233330121022-2010323033323323-0222220201013323-3032233310110131-0213110222313031-1012200230312322"></a>

## protected_cookies — protected_cookies / 222033330231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- protected_cookies

<a id="canonical-2020032312223312-0212212232031133-3330123312220112-3203300332221121-2002323120320111-0021333131321000-0113123012110322-0012332102112313"></a>

Type: `"list"`. Computed.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2311130201301132-1000103121330300-2331211321030012-2033021310311003-0232211120112001-0000023002331303-0203212300322120-1112202331210201"></a>

## Direct properties — protected_cookies / 222033330231 / 3

- [add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-3332312320111301-0221020103030022-3213101102123100-3210233313330331-1310222212231123-2130301330232322-2000013002022210-1131111301213003): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-3103032303213002-1002003012103201-0231331000200303-3120331023011112-3031121001221103-0030023221130123-1133133303122222-1202021030000322): complete subsection reference.

- [disable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332111131001023-1121113212020001-2010113210213021-0010233213310012-3310330322032331-0300301310211211-1013232233231221-3222301221030203): complete subsection reference.

- [enable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-2221011233302121-2002100113113302-0332000120300001-2213012233003032-1131133301221133-0132311310330303-2233112232013301-3032200323023221): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-3101113101300232-3030302222032113-2131003021320212-2203110210220320-3231020233100000-1012221032320110-1133331311232222-2110101132310203): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022011330020213-3203323031123113-2023123000131131-2202020303330303-2110333033302003-1223222312311102-0032101313211133-2233130113031233): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-023.md#canonical-1112303213003220-2233321130311001-1131213330200111-0123102102203121-0333223133122321-0022130233013103-2013120201121200-3000221310120323): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-023.md#canonical-2031220320331011-3213000230222112-2312333020302112-3012222132002322-0323121300222233-1101221101230121-3022010232320321-2233030112030211): complete subsection reference.

<a id="canonical-0220111010322223-1113300331111002-1210001113332021-3102130222002301-3123312331311231-1132222222010201-1021322102201003-1001322121212310"></a>

<a id="canonical-0333200130120313-0202203021012230-2112022213323231-3012211133002200-3110102003022031-3310012320231010-3233101020332210-1331002103333333"></a>

## max_age_value property — protected_cookies / 222033330231 / 4

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0322120232121001-2103022031003211-0333031023003312-2213012200020321-2120022311033332-2321313033323020-2002222212120330-2122132010110101"></a>

<a id="canonical-1332102130121030-3132200210213201-3323312301301321-3332011130001210-0210013120303111-1013001130120311-0033031301223123-0013012130331123"></a>

## name property — protected_cookies / 222033330231 / 5

Type: `"string"`. Computed.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-023.md#canonical-1123212322233202-0233321102200330-0002211133323011-2322000210121213-0322310000031221-3013130133232022-2232302100211102-1032223302003300): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-023.md#canonical-2120310030321121-3330023020102312-0032111100030010-3333301223020213-2231232123133030-1103110321010323-3113033123003312-3030122230310220): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-023.md#canonical-3023033032321212-1120133122002230-0202212130200011-1003023012130132-1011103132132210-2110023030320011-0233033111103022-3331323121301222): complete subsection reference.

<a id="canonical-1021221033213020-0010311303233011-2123130021332211-2113013122223001-2022322011221133-0113101132013333-3110310031310012-1301030212220011"></a>

## Next pages — protected_cookies / 222033330231 / 6

- [protected_cookies.add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-3332312320111301-0221020103030022-3213101102123100-3210233313330331-1310222212231123-2130301330232322-2000013002022210-1131111301213003)
- [protected_cookies.add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-3103032303213002-1002003012103201-0231331000200303-3120331023011112-3031121001221103-0030023221130123-1133133303122222-1202021030000322)
- [protected_cookies.disable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332111131001023-1121113212020001-2010113210213021-0010233213310012-3310330322032331-0300301310211211-1013232233231221-3222301221030203)
- [protected_cookies.enable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-2221011233302121-2002100113113302-0332000120300001-2213012233003032-1131133301221133-0132311310330303-2233112232013301-3032200323023221)
- [protected_cookies.ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-3101113101300232-3030302222032113-2131003021320212-2203110210220320-3231020233100000-1012221032320110-1133331311232222-2110101132310203)
- [protected_cookies.ignore_max_age](data-sources--http_loadbalancer--reference--group-022.md#canonical-3022011330020213-3203323031123113-2023123000131131-2202020303330303-2110333033302003-1223222312311102-0032101313211133-2233130113031233)
- [protected_cookies.ignore_samesite](data-sources--http_loadbalancer--reference--group-023.md#canonical-1112303213003220-2233321130311001-1131213330200111-0123102102203121-0333223133122321-0022130233013103-2013120201121200-3000221310120323)
- [protected_cookies.ignore_secure](data-sources--http_loadbalancer--reference--group-023.md#canonical-2031220320331011-3213000230222112-2312333020302112-3012222132002322-0323121300222233-1101221101230121-3022010232320321-2233030112030211)
- [protected_cookies.samesite_lax](data-sources--http_loadbalancer--reference--group-023.md#canonical-1123212322233202-0233321102200330-0002211133323011-2322000210121213-0322310000031221-3013130133232022-2232302100211102-1032223302003300)
- [protected_cookies.samesite_none](data-sources--http_loadbalancer--reference--group-023.md#canonical-2120310030321121-3330023020102312-0032111100030010-3333301223020213-2231232123133030-1103110321010323-3113033123003312-3030122230310220)
- [protected_cookies.samesite_strict](data-sources--http_loadbalancer--reference--group-023.md#canonical-3023033032321212-1120133122002230-0202212130200011-1003023012130132-1011103132132210-2110023030320011-0233033111103022-3331323121301222)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3332312320111301-0221020103030022-3213101102123100-3210233313330331-1310222212231123-2130301330232322-2000013002022210-1131111301213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311310301203033-2003132002310331-3203323011110013-0212110003312122-3003233122222102-0223312223033023-1232302002021202-0313111131010012"></a>

## protected_cookies.add_httponly — add_httponly / 202301321333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.add_httponly

<a id="canonical-0230021303020312-1002110112121301-0231133133023301-3203033011010100-2013000112111200-1302002101131100-0131202232010010-0321003301010330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1123113300310130-3312231030231120-3210212000031003-0220003000123302-2023012332003023-0020001033030321-2322203223030123-1331010223121233"></a>

## Direct properties — add_httponly / 202301321333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303213030022331-0330121003030131-0223110211121021-0333130021102013-2100022301123003-3300302130100033-2033111101303211-2032122132100213"></a>

## Next pages — add_httponly / 202301321333 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3103032303213002-1002003012103201-0231331000200303-3120331023011112-3031121001221103-0030023221130123-1133133303122222-1202021030000322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130002213112333-2201002200030020-0021020300223001-1112022203133332-2211133022103100-2322222310100210-2031310100132022-2203212113121102"></a>

## protected_cookies.add_secure — add_secure / 032232112102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.add_secure

<a id="canonical-2132121311203201-2003200010122201-1010321020200002-1313131002210032-0102312312111100-0112303020012112-0102120131302130-2200133012210113"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0223232100013211-3021310223103313-0121231302133233-2020322223223312-2220320022211312-2303310112301302-0221132200100222-1122020122300123"></a>

## Direct properties — add_secure / 032232112102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101323321003002-2032330223103100-0220032110022233-2322101322303222-2200300110321132-0001213101220000-2032021213010332-2211101001012031"></a>

## Next pages — add_secure / 032232112102 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2332111131001023-1121113212020001-2010113210213021-0010233213310012-3310330322032331-0300301310211211-1013232233231221-3222301221030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301111220233033-2012322013210000-2200012331131323-0031031103233111-0330002331223331-3312000130230011-3203333130031101-3212001310331323"></a>

## protected_cookies.disable_tampering_protection — disable_tampering_protection / 033003202230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.disable_tampering_protection

<a id="canonical-3310130223000132-1110002231113313-1213212213301320-2203013132112323-3213200233233222-2120030022031223-1220231231011010-1313232320021211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tampering protection.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2121330210020331-3302131331020220-3332001311203312-1331103110003121-0122101130012003-3201212210130303-3302311113113101-0320230320300001"></a>

## Direct properties — disable_tampering_protection / 033003202230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312321222320031-0030301201023220-0102210311101031-3030212022121012-2330102303002331-3031103200030102-3211033111111201-1310000132130131"></a>

## Next pages — disable_tampering_protection / 033003202230 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2221011233302121-2002100113113302-0332000120300001-2213012233003032-1131133301221133-0132311310330303-2233112232013301-3032200323023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331321103230300-2102110130202032-3132121030131232-1102022221101022-1112102322221310-3032002132212123-1202110001211013-0113113312003311"></a>

## protected_cookies.enable_tampering_protection — enable_tampering_protection / 330022201312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.enable_tampering_protection

<a id="canonical-0310030122012322-1100113010023333-3110220102103212-2113301110000221-0223130332132333-0320033020312021-1212221200320230-2110210012131003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tampering protection.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1310223231202112-2313232201000012-0233023211122011-2131130330023101-0100203303332303-2120111012000230-2120313322101101-0001013231310013"></a>

## Direct properties — enable_tampering_protection / 330022201312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331012000031133-2033032020003001-1313123223222020-1221300011212232-1101111020203011-0012121011010120-3301233322103120-3201031220312001"></a>

## Next pages — enable_tampering_protection / 330022201312 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3101113101300232-3030302222032113-2131003021320212-2203110210220320-3231020233100000-1012221032320110-1133331311232222-2110101132310203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033033322023001-2010222101102323-0113111100013121-2220002121320310-3203110330232132-2203332303211320-1301333020233323-0121103330222002"></a>

## protected_cookies.ignore_httponly — ignore_httponly / 221302213111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.ignore_httponly

<a id="canonical-3002032230032132-2331021121200000-3032211301012330-0323211220233102-1110011001023332-1211232333232332-3132030110203030-3001320201203023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1311030013123221-2103123030322210-0231211202202323-2333230102303030-3220332310203332-0230101013001120-2021020330230133-3013223012200213"></a>

## Direct properties — ignore_httponly / 221302213111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332230100213223-2210311330033213-3320231311312100-1131123310031123-3301110031220030-2323322003212222-0230210030131031-2231203330001201"></a>

## Next pages — ignore_httponly / 221302213111 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3022011330020213-3203323031123113-2023123000131131-2202020303330303-2110333033302003-1223222312311102-0032101313211133-2233130113031233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
