---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213310332012223-3301213230133330-0002122111011332-3233003303113301-1101120333202010-3331030332210312-1221021330023102-0301221111102231"></a>

## Property reference — Property reference / 000231010311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- Property reference

<a id="canonical-3112200201332232-0131023013113303-1130201031233030-0023022001301111-0303330232222220-2303120132202323-0120200020212000-1033023222323231"></a>

## Direct properties — Property reference / 000231010311 / 3

- [active_service_policies](data-sources--http_loadbalancer--reference--group-004.md#canonical-1030030122011102-1200203100110303-0222320330313000-1212333211001012-3011132203133320-3203212303220322-2213321232120023-2033013323323230): complete subsection reference.

<a id="canonical-3321022032313120-0313231222130313-2212223000003132-2112313332301103-2013002131231020-3020212322311132-0002302220021033-0210323111111122"></a>

<a id="canonical-3233311112112212-3113112013132203-2222123132223211-3213312031111202-1002303100221311-3031030222321130-2112220212232130-3330001331222001"></a>

## add_location property — Property reference / 000231010311 / 4

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

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

- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2230111322213330-1133313211131321-0000311132313220-1123102333032011-1231222323323010-0102210111000010-0332002211011211-1310131231110211): complete subsection reference.

- [advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2110110202303002-3212310231101322-0113133010033132-1231001030020312-3101021020302221-3211111113213332-0202113200011002-0213233111330230): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-0023021022222102-2330002011133123-3330210113010220-1322222203110001-2013001032123010-2211332013223321-1122030302003010-3011102323002033): complete subsection reference.

- [advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3233101001322201-3010101330010210-1230122302103210-1020231313123311-1123022210033012-0001211313213020-2123001023213012-0132232031212310): complete subsection reference.

<a id="canonical-1222201213311100-1033011230212231-0003030123011031-0011200123302030-2132220202033310-0310131330110203-3301200121222322-3000313303323312"></a>

<a id="canonical-2300012222320213-0310222123231133-1230221111011322-3302330133230213-1011100122032120-1022033231110331-0011103022031331-0323231332123130"></a>

## annotations property — Property reference / 000231010311 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120): complete subsection reference.

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311): complete subsection reference.

- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303): complete subsection reference.

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-3122013303110033-3021021311102313-3323021211213102-2100213133211231-2011301013301202-0233332230322311-0332101101001110-1223303130231312): complete subsection reference.

- [app_firewall](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202302022210302-0000322113313303-2232103000032320-1332112102002230-0330303303220310-3333030211302230-2332303332123021-2211321113113302): complete subsection reference.

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-3210121301002322-0000012333203131-2223303330303323-2311002110020111-3032020031220322-0311230032221020-1013000111213030-0032002000030003): complete subsection reference.

- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100): complete subsection reference.

- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131): complete subsection reference.

- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102): complete subsection reference.

- [captcha_challenge](data-sources--http_loadbalancer--reference--group-014.md#canonical-2333201102102232-2011010100211101-3321232300202320-0130113200321012-3222011121333011-3213310310031212-2203212330330331-3220013313301331): complete subsection reference.

- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233): complete subsection reference.

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323): complete subsection reference.

- [cors_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-1033333121011121-2330112221131202-0332012230130101-2330221123112012-1131023310033132-0021220023112313-1111023210121300-3213221010122332): complete subsection reference.

- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-3301321210130101-1003211310201200-2003003331000201-1010102230311020-3213103202230123-1233212032302110-1131123131102302-1031120113100113): complete subsection reference.

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003): complete subsection reference.

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001): complete subsection reference.

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331): complete subsection reference.

- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021): complete subsection reference.

- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123320333302212-3312111002110010-3032303202101023-0302100013110231-3132000001122200-0133332310312030-0321211001120321-3111232322003312): complete subsection reference.

- [default_sensitive_data_policy](data-sources--http_loadbalancer--reference--group-017.md#canonical-0111232021032210-2203300033211333-3203221201122011-1312201110201102-2011000032031020-1230113100330322-3230101122211000-2101102002000223): complete subsection reference.

<a id="canonical-0132333200200302-2113210001131010-0233231300023101-1330220310211132-3112101032302222-2220031110010003-1311201321333113-0300122303312023"></a>

<a id="canonical-2301003213313110-3030320020123013-1212200220321110-1310210131302301-1201003301022023-1123012313021121-0031122111311123-2203201302223302"></a>

## description property — Property reference / 000231010311 / 6

Type: `"string"`. Computed.

Description of the HTTPLoadBalancer.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [disable_api_definition](data-sources--http_loadbalancer--reference--group-017.md#canonical-3103320133122210-1101320300230032-2103233020111021-3331023123322321-0022223323012102-2001013112130132-3013002230133211-2332223011332022): complete subsection reference.

- [disable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-1302330231202101-0223003210100011-1103001000322103-2130013013301113-0031323333122311-1013313202011020-2010303010122300-2321321122232201): complete subsection reference.

- [disable_api_testing](data-sources--http_loadbalancer--reference--group-017.md#canonical-3333200003312023-1221031330131101-0030223302230000-0200202023223123-3230111313110123-3030133101330021-2320023222010202-1332313230012121): complete subsection reference.

- [disable_bot_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-0002113012001322-1331231230330100-0313110022323200-2010333332303100-2323110232011312-3300211312030310-0133233122112310-0233100023212113): complete subsection reference.

- [disable_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-0002302111031031-1130133033030120-1300022330230010-2000210133200112-3322203230333132-3220320102203131-0311010231031001-2120213121211103): complete subsection reference.

- [disable_client_side_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123332133130320-0331211322111000-2330232110321202-2330201211130101-3011112020212323-1221111213133101-3032301031112132-3223101322300031): complete subsection reference.

- [disable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-0112131023222001-3013102233003223-0212212313132321-3213111323012313-3103122301032310-1210111233031132-3121100100233010-0031322010212232): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-1330323130321022-1011312022300121-0130323013103312-0203031311203302-2120100101023303-2211110102033331-2312123223013030-3010231133332323): complete subsection reference.

- [disable_malware_protection](data-sources--http_loadbalancer--reference--group-017.md#canonical-2311133202211221-1301333010001102-3133200113211313-3133322202132313-0311131220320100-2120210322310223-1301000011201332-0001021230203033): complete subsection reference.

- [disable_rate_limit](data-sources--http_loadbalancer--reference--group-017.md#canonical-0220321111232132-0223122312102103-3110030000112122-2301210200110112-0032122230032323-2030101000021220-2131032321113013-2012300202230031): complete subsection reference.

- [disable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-1320232313112132-3030232230021301-1013212120230233-2020312000033331-1112112213013012-2121100232113103-2130221232312032-0300003230220011): complete subsection reference.

- [disable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-3013212331311311-1200011021032001-3302132230121201-1322113202300110-3312032230332123-2310111203203000-0113022033313023-0023311222003010): complete subsection reference.

- [disable_waf](data-sources--http_loadbalancer--reference--group-017.md#canonical-1213210303223130-0120021132231012-2321001110211123-3311211203301010-3202011221331212-0333122303302032-2211232123002030-1222233220133003): complete subsection reference.

- [do_not_advertise](data-sources--http_loadbalancer--reference--group-017.md#canonical-0312023203322213-3333031332332121-1301313200123121-1211123221033212-0103310121032331-0222122033112101-0230103313330001-3300110200032222): complete subsection reference.

<a id="canonical-2122001131201200-0133023301301111-3013032222131011-1201233110311131-1020233210000213-2220212022121222-1030300231120312-1011003332200210"></a>

<a id="canonical-3101301232002121-0033130020002013-3030100233323223-2030121022023223-2310300222331323-0022222131311111-0310013222133112-1312032221320310"></a>

## domains property — Property reference / 000231010311 / 7

Type: `["list", "string"]`. Computed.

List of Domains (host/authority header) that will be matched to load balancer. Supported Domains and
search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to load balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223): complete subsection reference.

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222): complete subsection reference.

- [enable_ip_reputation](data-sources--http_loadbalancer--reference--group-018.md#canonical-2222220232120233-1130313100210223-1110323001300223-2031010002311213-3032120333310102-3022103310121023-0103023331030110-0220211011320211): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-018.md#canonical-3112321013223011-0232120223301113-2123300000310001-2303013212022230-1020121313200102-3113021131130312-3132303103233220-0130203210202122): complete subsection reference.

- [enable_threat_mesh](data-sources--http_loadbalancer--reference--group-018.md#canonical-1111231022131133-1331313200113310-3310302011012003-0030332010313023-0230102231302212-2013001002111223-2120221113320103-0133233131213321): complete subsection reference.

- [enable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223302231131311-1212022331012012-1201212100221020-3222300100201333-2302211010002001-3333320223120100-2001323010222333-0112102133220100): complete subsection reference.

- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232): complete subsection reference.

- [http](data-sources--http_loadbalancer--reference--group-018.md#canonical-1221000122101103-3100213033303232-2132010133302010-2012231203222223-2220001112313332-0002130322200033-3133223330023000-2102003233113221): complete subsection reference.

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301): complete subsection reference.

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131): complete subsection reference.

<a id="canonical-3311013102232313-2212132200012202-1301123130323103-1330101223003211-3323011200202312-2320210223102322-2213021211020320-2012200011320233"></a>

<a id="canonical-0333033220333021-3202031311300113-2320110123222032-0111231330211100-0131112012312311-0330221023211021-0221312301022203-2331310030232100"></a>

## ID property — Property reference / 000231010311 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-3322121320310111-0021311031133102-3002213332023220-1023112102022302-3222311022023111-1222120222313303-1310032321220233-2003030001231111): complete subsection reference.

- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130): complete subsection reference.

- [l7_ddos_action_block](data-sources--http_loadbalancer--reference--group-020.md#canonical-3121023203221011-0201101021002021-0101212311232312-2121320103303333-0123032110032312-3313111220330000-2332202302121131-2130131301133210): complete subsection reference.

- [l7_ddos_action_default](data-sources--http_loadbalancer--reference--group-020.md#canonical-1321021012132031-3001312230330301-0022003310222320-2211111231213322-2010212203031011-2021010223131312-3020200133123132-2023213310130210): complete subsection reference.

- [l7_ddos_action_js_challenge](data-sources--http_loadbalancer--reference--group-020.md#canonical-3301313020211313-2011033021131302-2123320013103200-0212212231232132-2313220232122330-1120020310001123-3110221330102102-2001223123013231): complete subsection reference.

- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-020.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133): complete subsection reference.

<a id="canonical-3302332000203231-2310231120013221-2022032331310001-2001213231230121-1012100031202302-3323310022112121-1310110231322320-2220033111122311"></a>

<a id="canonical-2222202110300310-3330200031003323-1323022113312332-0223012013131331-2100011333020002-3001313010233301-1032011010310222-3202230312313112"></a>

## labels property — Property reference / 000231010311 / 9

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [least_active](data-sources--http_loadbalancer--reference--group-020.md#canonical-3021012323102120-3202203000032032-2133022123330310-0310112201022300-0323033113303211-2210322220030333-2103012033230122-3232312030230201): complete subsection reference.

- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113): complete subsection reference.

- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101): complete subsection reference.

- [multi_lb_app](data-sources--http_loadbalancer--reference--group-021.md#canonical-3312033113300312-1002132313113311-3233122303230311-2103222022222003-3000233023022122-2023233330322332-1333021223220210-1021232333203310): complete subsection reference.

<a id="canonical-3032332313213233-0020222012030312-3301300012303032-3102033021310032-3330221210210020-0310101303101230-0120310301300012-1303211100233311"></a>

<a id="canonical-3211032220200230-2303013311010303-3222022110100312-2133233020003210-0011021320122131-2032030223023203-3022213133311012-2120121121322020"></a>

## name property — Property reference / 000231010311 / 10

Type: `"string"`. Required.

Name of the HTTPLoadBalancer.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2033100103103132-1002321111221031-0211123202320112-3110320221033100-2303113232102323-3003203103121002-2320110222112113-3030013302121130"></a>

<a id="canonical-1022022000021002-0230223102221330-1033113330103301-2103122201223132-3303013102303033-1011001023100320-1132122303233231-3302132200330020"></a>

## namespace property — Property reference / 000231010311 / 11

Type: `"string"`. Required.

Namespace where the HTTPLoadBalancer exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-0033322013033333-1012021230311113-3030032222123110-2130202012323311-2130011103202221-3200200101001031-2230202203012132-2012120233023302): complete subsection reference.

- [no_service_policies](data-sources--http_loadbalancer--reference--group-021.md#canonical-1012220231323000-0002101033211030-3111020332233130-0011331233023201-2102321003221301-1230013230021133-3230200102213220-0022221221200321): complete subsection reference.

- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312): complete subsection reference.

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310): complete subsection reference.

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111): complete subsection reference.

- [random](data-sources--http_loadbalancer--reference--group-023.md#canonical-3211233301202123-3021122113200023-3121032321113322-3231020200122333-1021331133012303-3330210102330302-1000033231000302-3302120223110221): complete subsection reference.

- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321): complete subsection reference.

- [ring_hash](data-sources--http_loadbalancer--reference--group-023.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323): complete subsection reference.

- [round_robin](data-sources--http_loadbalancer--reference--group-023.md#canonical-2100021222003130-0100011231321302-3110002003123302-1203123023020200-3232130013121212-2002101032122031-3302132222233331-3002022203023111): complete subsection reference.

- [routes](data-sources--http_loadbalancer--reference--group-023.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000): complete subsection reference.

- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301): complete subsection reference.

- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-1111113201333121-3123100001123201-0011013100212320-1131120220102201-2312121211133221-0310211131231233-2223203221032301-2301031302111331): complete subsection reference.

- [service_policies_from_namespace](data-sources--http_loadbalancer--reference--group-026.md#canonical-3032010102013301-1011233111333022-1301232320231300-1133103120030100-2321013303210022-2213102120021010-1000322312230310-2230332300333132): complete subsection reference.

- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310): complete subsection reference.

- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-026.md#canonical-1123002022123300-1112121112130100-2221312020130023-3123011020012032-0303121031103103-2302330113032012-2232201101111100-3332313211000001): complete subsection reference.

- [source_ip_stickiness](data-sources--http_loadbalancer--reference--group-026.md#canonical-1132130232001202-3132322001023103-2220102200330221-3102023123230101-2011032011032233-0011310023211220-1012121130020101-2232213232312220): complete subsection reference.

- [system_default_timeouts](data-sources--http_loadbalancer--reference--group-026.md#canonical-1002113332100220-0210021202232223-2330212033110033-3003023131320231-2021130022231113-0331333211332210-3013003301000123-3121110021322331): complete subsection reference.

- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020): complete subsection reference.

- [user_id_client_ip](data-sources--http_loadbalancer--reference--group-026.md#canonical-2223203020310031-1320313232132320-1010131100213310-3102302311120331-1313020012310011-2302030032233032-3322311032002211-0310020323003300): complete subsection reference.

- [user_identification](data-sources--http_loadbalancer--reference--group-026.md#canonical-3300213310030023-1303230033010211-3001312312000330-2020213133231032-1301233011000133-1301323210201312-0220100111121100-0231103130020022): complete subsection reference.

- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332): complete subsection reference.
