# Changelog

<a name="v1.7.2-cap-CR-26082"></a>
## [v1.7.2-cap-CR-26082](https://github.com/argoproj/argo-rollouts/compare/v1.7.1-CR-24605...v1.7.2-cap-CR-26082) (2025-01-30)

* fix: avoid using root user in plugin container (#1256)

* update version to Codefresh variant
* security bump for https://osv.dev/vulnerability/GHSA-82m2-cv7p-4m75
* update codegen tools
* security fix https://nvd.nist.gov/vuln/detail/CVE-2023-46402
* **deps:** bump google.golang.org/grpc from 1.62.1 to 1.63.0 ([#3497](https://github.com/argoproj/argo-rollouts/issues/3497))

### Docs

* GitHub notification reference

### Fix

* replicaSet not scaled down due to incorrect annotations ([#3762](https://github.com/argoproj/argo-rollouts/issues/3762)) ([#3784](https://github.com/argoproj/argo-rollouts/issues/3784))
* add update verb to ClusterRole permissions for scaleDown feature. Fixes [#3672](https://github.com/argoproj/argo-rollouts/issues/3672) ([#3675](https://github.com/argoproj/argo-rollouts/issues/3675))
* **analysis:** explicitly set datadog aggregator to last only on v2 ([#3730](https://github.com/argoproj/argo-rollouts/issues/3730))
* **analysis:** Take RollbackWindow into account when Reconciling Analysis Runs. Fixes [#3669](https://github.com/argoproj/argo-rollouts/issues/3669) ([#3670](https://github.com/argoproj/argo-rollouts/issues/3670))
* **controller:** Get the right resourceName for traefik.io.Fixes [#3615](https://github.com/argoproj/argo-rollouts/issues/3615) ([#3759](https://github.com/argoproj/argo-rollouts/issues/3759))
* **controller:** use the stableRS from the rollout context rather tha… ([#3664](https://github.com/argoproj/argo-rollouts/issues/3664))
* **dashboard:** Update pod status logic to support native sidecars. Fixes [#3366](https://github.com/argoproj/argo-rollouts/issues/3366) ([#3639](https://github.com/argoproj/argo-rollouts/issues/3639))
* **metricprovider:** reuse http.Transport for http.Client ([#3780](https://github.com/argoproj/argo-rollouts/issues/3780))


<a name="v1.7.1-CR-24605"></a>
## [v1.7.1-CR-24605](https://github.com/argoproj/argo-rollouts/compare/1.6.1...v1.7.1-CR-24605) (2024-07-26)

### Build

* **deps:** always resolve momentjs version 2.29.4 ([#3182](https://github.com/argoproj/argo-rollouts/issues/3182))

### Chore

* fix PodSecurity warning ([#3424](https://github.com/argoproj/argo-rollouts/issues/3424))
* add WeLab Bank to users.md ([#2996](https://github.com/argoproj/argo-rollouts/issues/2996))
* change file name for readthedocs compatibility ([#2999](https://github.com/argoproj/argo-rollouts/issues/2999))
* Update users doc with CircleCI ([#3028](https://github.com/argoproj/argo-rollouts/issues/3028))
* updating getCanaryConfigId to be more efficient with better error handling ([#3070](https://github.com/argoproj/argo-rollouts/issues/3070))
* bump k8s versions to 1.29 ([#3494](https://github.com/argoproj/argo-rollouts/issues/3494))
* add missing rollout fields ([#3062](https://github.com/argoproj/argo-rollouts/issues/3062))
* upgrade cosign ([#3139](https://github.com/argoproj/argo-rollouts/issues/3139))
* add OpenSSF Scorecard badge ([#3154](https://github.com/argoproj/argo-rollouts/issues/3154))
* add test for reconcileEphemeralMetadata() ([#3163](https://github.com/argoproj/argo-rollouts/issues/3163))
* Add exception to `requireCanaryStableServices` to disable validation when using the `hashicorp/consul` plugin ([#3339](https://github.com/argoproj/argo-rollouts/issues/3339))
* leave the validation of setHeaderRoute to the plugin when plugins is not empty. ([#2898](https://github.com/argoproj/argo-rollouts/issues/2898))
* fix lint errors reported by golangci-lint ([#3458](https://github.com/argoproj/argo-rollouts/issues/3458))
* fix unit test data races ([#3478](https://github.com/argoproj/argo-rollouts/issues/3478)) ([#3479](https://github.com/argoproj/argo-rollouts/issues/3479))
* added organization to users.md ([#3481](https://github.com/argoproj/argo-rollouts/issues/3481))
* set webpack hashFunction to modern sha256, remove legacy-provider. Fixes [#2609](https://github.com/argoproj/argo-rollouts/issues/2609) ([#3475](https://github.com/argoproj/argo-rollouts/issues/3475))
* remove year from codegen license  ([#3282](https://github.com/argoproj/argo-rollouts/issues/3282))
* update follow-redirects to 1.15.5 ([#3314](https://github.com/argoproj/argo-rollouts/issues/3314))
* add logging context around replicaset updates ([#3326](https://github.com/argoproj/argo-rollouts/issues/3326))
* bump notification engine lib ([#3327](https://github.com/argoproj/argo-rollouts/issues/3327))
* change controller's deploy strategy to RollingUpdate due to leader election ([#3334](https://github.com/argoproj/argo-rollouts/issues/3334))
* leave the validation of setHeaderRoute to the plugin when plugins is not empty. ([#2898](https://github.com/argoproj/argo-rollouts/issues/2898))
* Update notifications engine to 7a06976 ([#3384](https://github.com/argoproj/argo-rollouts/issues/3384))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.25.11 to 1.25.12 ([#3230](https://github.com/argoproj/argo-rollouts/issues/3230))
* **deps:** bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#3422](https://github.com/argoproj/argo-rollouts/issues/3422))
* **deps:** bump softprops/action-gh-release from 2.0.3 to 2.0.4 ([#3442](https://github.com/argoproj/argo-rollouts/issues/3442))
* **deps:** bump softprops/action-gh-release from 2.0.2 to 2.0.3 ([#3440](https://github.com/argoproj/argo-rollouts/issues/3440))
* **deps:** bump softprops/action-gh-release from 1 to 2 ([#3438](https://github.com/argoproj/argo-rollouts/issues/3438))
* **deps:** bump docker/build-push-action from 5.1.0 to 5.2.0 ([#3439](https://github.com/argoproj/argo-rollouts/issues/3439))
* **deps:** bump docker/setup-buildx-action from 3.1.0 to 3.2.0 ([#3449](https://github.com/argoproj/argo-rollouts/issues/3449))
* **deps:** bump google.golang.org/grpc from 1.62.0 to 1.62.1 ([#3426](https://github.com/argoproj/argo-rollouts/issues/3426))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.27.4 to 1.27.5 ([#3421](https://github.com/argoproj/argo-rollouts/issues/3421))
* **deps:** bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#3419](https://github.com/argoproj/argo-rollouts/issues/3419))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.27.0 to 1.27.4 ([#3410](https://github.com/argoproj/argo-rollouts/issues/3410))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.27.0 to 1.30.1 ([#3399](https://github.com/argoproj/argo-rollouts/issues/3399))
* **deps:** bump google.golang.org/grpc from 1.61.0 to 1.62.0 ([#3404](https://github.com/argoproj/argo-rollouts/issues/3404))
* **deps:** bump docker/setup-buildx-action from 3.0.0 to 3.1.0 ([#3406](https://github.com/argoproj/argo-rollouts/issues/3406))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.33.0 to 1.36.1 ([#3400](https://github.com/argoproj/argo-rollouts/issues/3400))
* **deps:** bump codecov/codecov-action from 4.0.1 to 4.1.0 ([#3403](https://github.com/argoproj/argo-rollouts/issues/3403))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.30.1 to 1.30.3 ([#3447](https://github.com/argoproj/argo-rollouts/issues/3447))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.6 to 1.27.0 ([#3368](https://github.com/argoproj/argo-rollouts/issues/3368))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.32.2 to 1.33.0 ([#3363](https://github.com/argoproj/argo-rollouts/issues/3363))
* **deps:** bump docker/login-action from 3.0.0 to 3.1.0 ([#3443](https://github.com/argoproj/argo-rollouts/issues/3443))
* **deps:** bump golang.org/x/oauth2 from 0.16.0 to 0.17.0 ([#3357](https://github.com/argoproj/argo-rollouts/issues/3357))
* **deps:** bump golangci/golangci-lint-action from 3 to 4 ([#3359](https://github.com/argoproj/argo-rollouts/issues/3359))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.7 to 1.27.0 ([#3341](https://github.com/argoproj/argo-rollouts/issues/3341))
* **deps:** bump peter-evans/create-pull-request from 5 to 6 ([#3342](https://github.com/argoproj/argo-rollouts/issues/3342))
* **deps:** bump sigstore/cosign-installer from 3.3.0 to 3.4.0 ([#3343](https://github.com/argoproj/argo-rollouts/issues/3343))
* **deps:** bump codecov/codecov-action from 3.1.5 to 4.0.1 ([#3347](https://github.com/argoproj/argo-rollouts/issues/3347))
* **deps:** bump github.com/evanphx/json-patch/v5 from 5.8.1 to 5.9.0 ([#3335](https://github.com/argoproj/argo-rollouts/issues/3335))
* **deps:** bump docker/build-push-action from 5.2.0 to 5.3.0 ([#3448](https://github.com/argoproj/argo-rollouts/issues/3448))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.5 to 1.26.6 ([#3322](https://github.com/argoproj/argo-rollouts/issues/3322))
* **deps:** bump github.com/evanphx/json-patch/v5 from 5.8.0 to 5.8.1 ([#3312](https://github.com/argoproj/argo-rollouts/issues/3312))
* **deps:** bump codecov/codecov-action from 3.1.4 to 3.1.5 ([#3330](https://github.com/argoproj/argo-rollouts/issues/3330))
* **deps:** bump slsa-framework/slsa-github-generator from 1.9.0 to 1.9.1 ([#3456](https://github.com/argoproj/argo-rollouts/issues/3456))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.36.1 to 1.36.3 ([#3452](https://github.com/argoproj/argo-rollouts/issues/3452))
* **deps:** bump google.golang.org/grpc from 1.60.1 to 1.61.0 ([#3325](https://github.com/argoproj/argo-rollouts/issues/3325))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.4 to 1.26.5 ([#3319](https://github.com/argoproj/argo-rollouts/issues/3319))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.3 to 1.26.4 ([#3313](https://github.com/argoproj/argo-rollouts/issues/3313))
* **deps:** bump actions/cache from 3 to 4 ([#3315](https://github.com/argoproj/argo-rollouts/issues/3315))
* **deps:** bump slsa-framework/slsa-github-generator from 1.9.1 to 1.10.0 ([#3462](https://github.com/argoproj/argo-rollouts/issues/3462))
* **deps:** bump github.com/evanphx/json-patch/v5 from 5.7.0 to 5.8.0 ([#3309](https://github.com/argoproj/argo-rollouts/issues/3309))
* **deps:** bump golang.org/x/oauth2 from 0.15.0 to 0.16.0 ([#3294](https://github.com/argoproj/argo-rollouts/issues/3294))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.32.1 to 1.32.2 ([#3288](https://github.com/argoproj/argo-rollouts/issues/3288))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.2 to 1.26.3 ([#3289](https://github.com/argoproj/argo-rollouts/issues/3289))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.6 to 1.26.7 ([#3290](https://github.com/argoproj/argo-rollouts/issues/3290))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.24.0 to 1.24.1 ([#3291](https://github.com/argoproj/argo-rollouts/issues/3291))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.30.3 to 1.30.4 ([#3461](https://github.com/argoproj/argo-rollouts/issues/3461))
* **deps:** bump google.golang.org/protobuf from 1.31.0 to 1.32.0 ([#3273](https://github.com/argoproj/argo-rollouts/issues/3273))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.1 to 1.26.2 ([#3268](https://github.com/argoproj/argo-rollouts/issues/3268))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.5 to 1.26.6 ([#3269](https://github.com/argoproj/argo-rollouts/issues/3269))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.32.0 to 1.32.1 ([#3270](https://github.com/argoproj/argo-rollouts/issues/3270))
* **deps:** bump google.golang.org/grpc from 1.60.0 to 1.60.1 ([#3260](https://github.com/argoproj/argo-rollouts/issues/3260))
* **deps:** bump github/codeql-action from 2 to 3 ([#3252](https://github.com/argoproj/argo-rollouts/issues/3252))
* **deps:** bump actions/upload-artifact from 3 to 4 ([#3255](https://github.com/argoproj/argo-rollouts/issues/3255))
* **deps:** bump sigstore/cosign-installer from 3.2.0 to 3.3.0 ([#3245](https://github.com/argoproj/argo-rollouts/issues/3245))
* **deps:** bump google.golang.org/grpc from 1.59.0 to 1.60.0 ([#3246](https://github.com/argoproj/argo-rollouts/issues/3246))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.26.0 to 1.26.1 ([#3241](https://github.com/argoproj/argo-rollouts/issues/3241))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.4 to 1.26.5 ([#3240](https://github.com/argoproj/argo-rollouts/issues/3240))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.31.4 to 1.32.0 ([#3239](https://github.com/argoproj/argo-rollouts/issues/3239))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.25.12 to 1.26.0 ([#3236](https://github.com/argoproj/argo-rollouts/issues/3236))
* **deps:** bump codecov/codecov-action from 4.1.0 to 4.1.1 ([#3476](https://github.com/argoproj/argo-rollouts/issues/3476))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 from 2.12.4 to 2.13.0 ([#3217](https://github.com/argoproj/argo-rollouts/issues/3217))
* **deps:** bump actions/stale from 8 to 9 ([#3232](https://github.com/argoproj/argo-rollouts/issues/3232))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.31.3 to 1.31.4 ([#3235](https://github.com/argoproj/argo-rollouts/issues/3235))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.3 to 1.26.4 ([#3234](https://github.com/argoproj/argo-rollouts/issues/3234))
* **deps:** update golang to 1.21 ([#3482](https://github.com/argoproj/argo-rollouts/issues/3482))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.31.2 to 1.31.3 ([#3226](https://github.com/argoproj/argo-rollouts/issues/3226))
* **deps:** bump actions/setup-python from 4 to 5 ([#3227](https://github.com/argoproj/argo-rollouts/issues/3227))
* **deps:** bump actions/setup-go from 4.1.0 to 5.0.0 ([#3228](https://github.com/argoproj/argo-rollouts/issues/3228))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.2 to 1.26.3 ([#3229](https://github.com/argoproj/argo-rollouts/issues/3229))
* **deps:** Bump k8s dependencies to v1.26.11 ([#3211](https://github.com/argoproj/argo-rollouts/issues/3211))
* **deps:** bump argo-ui and fix browser console errors ([#3212](https://github.com/argoproj/argo-rollouts/issues/3212))
* **deps:** bump docker/build-push-action from 5.0.0 to 5.1.0 ([#3178](https://github.com/argoproj/argo-rollouts/issues/3178))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.25.10 to 1.25.11 ([#3206](https://github.com/argoproj/argo-rollouts/issues/3206))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.26.1 to 1.26.2 ([#3207](https://github.com/argoproj/argo-rollouts/issues/3207))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.31.1 to 1.31.2 ([#3208](https://github.com/argoproj/argo-rollouts/issues/3208))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.30.5 to 1.31.1 ([#3201](https://github.com/argoproj/argo-rollouts/issues/3201))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.25.2 to 1.26.1 ([#3203](https://github.com/argoproj/argo-rollouts/issues/3203))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.25.8 to 1.25.10 ([#3204](https://github.com/argoproj/argo-rollouts/issues/3204))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.25.5 to 1.25.8 ([#3191](https://github.com/argoproj/argo-rollouts/issues/3191))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.24.3 to 1.25.2 ([#3192](https://github.com/argoproj/argo-rollouts/issues/3192))
* **deps:** bump golang.org/x/oauth2 from 0.13.0 to 0.15.0 ([#3187](https://github.com/argoproj/argo-rollouts/issues/3187))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.30.3 to 1.30.5 ([#3193](https://github.com/argoproj/argo-rollouts/issues/3193))
* **deps:** bump github.com/antonmedv/expr from 1.15.4 to 1.15.5 ([#3186](https://github.com/argoproj/argo-rollouts/issues/3186))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.30.1 to 1.30.3 ([#3179](https://github.com/argoproj/argo-rollouts/issues/3179))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.24.0 to 1.24.3 ([#3180](https://github.com/argoproj/argo-rollouts/issues/3180))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 from 2.12.3 to 2.12.4 ([#3150](https://github.com/argoproj/argo-rollouts/issues/3150))
* **deps:** bump github.com/antonmedv/expr from 1.15.3 to 1.15.4 ([#3184](https://github.com/argoproj/argo-rollouts/issues/3184))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.23.0 to 1.25.5 ([#3183](https://github.com/argoproj/argo-rollouts/issues/3183))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.30.0 to 1.30.1 ([#3166](https://github.com/argoproj/argo-rollouts/issues/3166))
* **deps:** bump github.com/hashicorp/go-plugin from 1.5.2 to 1.6.0 ([#3167](https://github.com/argoproj/argo-rollouts/issues/3167))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.27.5 to 1.27.9 ([#3469](https://github.com/argoproj/argo-rollouts/issues/3469))
* **deps:** bump github.com/bombsimon/logrusr/v4 from 4.0.0 to 4.1.0 ([#3151](https://github.com/argoproj/argo-rollouts/issues/3151))
* **deps:** bump github.com/spf13/cobra from 1.7.0 to 1.8.0 ([#3152](https://github.com/argoproj/argo-rollouts/issues/3152))
* **deps:** bump sigstore/cosign-installer from 3.1.2 to 3.2.0 ([#3158](https://github.com/argoproj/argo-rollouts/issues/3158))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.22.0 to 1.23.0 ([#3161](https://github.com/argoproj/argo-rollouts/issues/3161))
* **deps:** bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#3429](https://github.com/argoproj/argo-rollouts/issues/3429))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.28.0 to 1.30.0 ([#3144](https://github.com/argoproj/argo-rollouts/issues/3144))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.22.0 to 1.24.0 ([#3143](https://github.com/argoproj/argo-rollouts/issues/3143))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.20.0 to 1.22.0 ([#3149](https://github.com/argoproj/argo-rollouts/issues/3149))
* **deps:** bump github.com/aws/smithy-go from 1.20.1 to 1.20.2 ([#3488](https://github.com/argoproj/argo-rollouts/issues/3488))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.19.1 to 1.20.0 ([#3135](https://github.com/argoproj/argo-rollouts/issues/3135))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.21.2 to 1.22.0 ([#3136](https://github.com/argoproj/argo-rollouts/issues/3136))
* **deps:** bump sigs.k8s.io/yaml from 1.3.0 to 1.4.0 ([#3122](https://github.com/argoproj/argo-rollouts/issues/3122))
* **deps:** bump google.golang.org/grpc from 1.58.3 to 1.59.0 ([#3113](https://github.com/argoproj/argo-rollouts/issues/3113))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.21.6 to 1.22.0 ([#3127](https://github.com/argoproj/argo-rollouts/issues/3127))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.19.0 to 1.19.1 ([#3123](https://github.com/argoproj/argo-rollouts/issues/3123))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.9 to 1.28.0 ([#3124](https://github.com/argoproj/argo-rollouts/issues/3124))
* **deps:** bump golang.org/x/oauth2 from 0.10.0 to 0.13.0 ([#3107](https://github.com/argoproj/argo-rollouts/issues/3107))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.45 to 1.19.0 ([#3109](https://github.com/argoproj/argo-rollouts/issues/3109))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.44 to 1.18.45 ([#3101](https://github.com/argoproj/argo-rollouts/issues/3101))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.21.4 to 1.21.6 ([#3100](https://github.com/argoproj/argo-rollouts/issues/3100))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.8 to 1.27.9 ([#3102](https://github.com/argoproj/argo-rollouts/issues/3102))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.21.1 to 1.21.2 ([#3103](https://github.com/argoproj/argo-rollouts/issues/3103))
* **deps:** bump github.com/prometheus/common from 0.42.0 to 0.51.1 ([#3468](https://github.com/argoproj/argo-rollouts/issues/3468))
* **deps:** bump google.golang.org/grpc from 1.58.2 to 1.58.3 ([#3098](https://github.com/argoproj/argo-rollouts/issues/3098))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.43 to 1.18.44 ([#3099](https://github.com/argoproj/argo-rollouts/issues/3099))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.21.0 to 1.21.1 ([#3085](https://github.com/argoproj/argo-rollouts/issues/3085))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.7 to 1.27.8 ([#3086](https://github.com/argoproj/argo-rollouts/issues/3086))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.42 to 1.18.43 ([#3072](https://github.com/argoproj/argo-rollouts/issues/3072))
* **deps:** bump github.com/hashicorp/go-plugin from 1.5.1 to 1.5.2 ([#3056](https://github.com/argoproj/argo-rollouts/issues/3056))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.26.0 to 1.26.1 ([#3490](https://github.com/argoproj/argo-rollouts/issues/3490))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.41 to 1.18.42 ([#3055](https://github.com/argoproj/argo-rollouts/issues/3055))
* **deps:** bump github.com/antonmedv/expr from 1.15.2 to 1.15.3 ([#3046](https://github.com/argoproj/argo-rollouts/issues/3046))
* **deps:** bump docker/setup-qemu-action from 2.2.0 to 3.0.0 ([#3031](https://github.com/argoproj/argo-rollouts/issues/3031))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.39 to 1.18.41 ([#3047](https://github.com/argoproj/argo-rollouts/issues/3047))
* **deps:** bump google.golang.org/grpc from 1.58.0 to 1.58.2 ([#3050](https://github.com/argoproj/argo-rollouts/issues/3050))
* **deps:** bump google.golang.org/grpc from 1.57.0 to 1.58.0 ([#3023](https://github.com/argoproj/argo-rollouts/issues/3023))
* **deps:** bump github.com/evanphx/json-patch/v5 from 5.6.0 to 5.7.0 ([#3030](https://github.com/argoproj/argo-rollouts/issues/3030))
* **deps:** bump docker/metadata-action from 4 to 5 ([#3032](https://github.com/argoproj/argo-rollouts/issues/3032))
* **deps:** bump docker/build-push-action from 4.1.1 to 5.0.0 ([#3033](https://github.com/argoproj/argo-rollouts/issues/3033))
* **deps:** bump docker/setup-buildx-action from 2.10.0 to 3.0.0 ([#3034](https://github.com/argoproj/argo-rollouts/issues/3034))
* **deps:** bump docker/login-action from 2.2.0 to 3.0.0 ([#3035](https://github.com/argoproj/argo-rollouts/issues/3035))
* **deps:** bump github.com/antonmedv/expr from 1.15.1 to 1.15.2 ([#3036](https://github.com/argoproj/argo-rollouts/issues/3036))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.36.3 to 1.37.0 ([#3489](https://github.com/argoproj/argo-rollouts/issues/3489))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.21.3 to 1.21.4 ([#3025](https://github.com/argoproj/argo-rollouts/issues/3025))
* **deps:** bump github.com/hashicorp/go-plugin from 1.5.0 to 1.5.1 ([#3017](https://github.com/argoproj/argo-rollouts/issues/3017))
* **deps:** bump github.com/antonmedv/expr from 1.13.0 to 1.15.1 ([#3024](https://github.com/argoproj/argo-rollouts/issues/3024))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.38 to 1.18.39 ([#3018](https://github.com/argoproj/argo-rollouts/issues/3018))
* **deps:** bump actions/checkout from 3 to 4 ([#3012](https://github.com/argoproj/argo-rollouts/issues/3012))
* **deps:** bump sigstore/cosign-installer from 3.1.1 to 3.1.2 ([#3011](https://github.com/argoproj/argo-rollouts/issues/3011))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.37 to 1.18.38 ([#3002](https://github.com/argoproj/argo-rollouts/issues/3002))
* **deps:** bump github.com/hashicorp/go-plugin from 1.4.10 to 1.5.0 ([#2995](https://github.com/argoproj/argo-rollouts/issues/2995))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.27.9 to 1.27.10 ([#3492](https://github.com/argoproj/argo-rollouts/issues/3492))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.30.4 to 1.30.5 ([#3491](https://github.com/argoproj/argo-rollouts/issues/3491))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.6 to 1.27.7 ([#2990](https://github.com/argoproj/argo-rollouts/issues/2990))
* **deps:** bump docker/setup-buildx-action from 2.9.1 to 2.10.0 ([#2994](https://github.com/argoproj/argo-rollouts/issues/2994))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.21.0 to 1.21.3 ([#2977](https://github.com/argoproj/argo-rollouts/issues/2977))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.36 to 1.18.37 ([#2984](https://github.com/argoproj/argo-rollouts/issues/2984))
* **deps:** bump slsa-framework/slsa-github-generator from 1.8.0 to 1.9.0 ([#2983](https://github.com/argoproj/argo-rollouts/issues/2983))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.33 to 1.18.36 ([#2978](https://github.com/argoproj/argo-rollouts/issues/2978))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.2 to 1.27.6 ([#2979](https://github.com/argoproj/argo-rollouts/issues/2979))

### Docs

* more best practices ([#3484](https://github.com/argoproj/argo-rollouts/issues/3484))
* typo in BlueGreen ([#3463](https://github.com/argoproj/argo-rollouts/issues/3463))
* minor readability on migration ([#3427](https://github.com/argoproj/argo-rollouts/issues/3427))
* added Consul plugin support to website ([#3362](https://github.com/argoproj/argo-rollouts/issues/3362))
* Update shell autocompletion instructions ([#3377](https://github.com/argoproj/argo-rollouts/issues/3377))
* Update Changelog ([#3365](https://github.com/argoproj/argo-rollouts/issues/3365))
* Guides for popular use-cases ([#3346](https://github.com/argoproj/argo-rollouts/issues/3346))
* Update Changelog ([#3328](https://github.com/argoproj/argo-rollouts/issues/3328))
* Fixed the key for headers in prometheus based argo analysis ([#3306](https://github.com/argoproj/argo-rollouts/issues/3306))
* mention archival of the SMI spec ([#3263](https://github.com/argoproj/argo-rollouts/issues/3263))
* Update Changelog ([#3244](https://github.com/argoproj/argo-rollouts/issues/3244))
* Update Changelog ([#3214](https://github.com/argoproj/argo-rollouts/issues/3214))
* Update Changelog ([#2952](https://github.com/argoproj/argo-rollouts/issues/2952))
* fix typo in smi.md ([#3160](https://github.com/argoproj/argo-rollouts/issues/3160))
* Update Changelog ([#3148](https://github.com/argoproj/argo-rollouts/issues/3148))
* add Gateway-API integration information to README.md ([#2985](https://github.com/argoproj/argo-rollouts/issues/2985))
* add CONTRIBUTING.md at root of repo, directing to docs/ ([#3121](https://github.com/argoproj/argo-rollouts/issues/3121))
* Ensure image not present between incomplete sentence. ([#3079](https://github.com/argoproj/argo-rollouts/issues/3079))
* clarify external clusters ([#3058](https://github.com/argoproj/argo-rollouts/issues/3058))
* Update Changelog ([#3021](https://github.com/argoproj/argo-rollouts/issues/3021))
* replace `patchesStrategicMerge` with `patches` in tests/docs ([#3010](https://github.com/argoproj/argo-rollouts/issues/3010))
* update all ingress objects to networking.k8s.io/v1 ([#3005](https://github.com/argoproj/argo-rollouts/issues/3005))
* Remove rogue apostrophe in features/analysis.md ([#3001](https://github.com/argoproj/argo-rollouts/issues/3001))
* add contour integration information to README.md ([#2980](https://github.com/argoproj/argo-rollouts/issues/2980))
* **analysis:** Add note about availability of new datadog v2 functionality ([#3131](https://github.com/argoproj/argo-rollouts/issues/3131))
* **deps:** Specify minimum kustomize version ([#3199](https://github.com/argoproj/argo-rollouts/issues/3199))

### Feat

* Support OAuth2 for prometheus and web providers ([#3038](https://github.com/argoproj/argo-rollouts/issues/3038))
* add command args for plugin ([#2992](https://github.com/argoproj/argo-rollouts/issues/2992))
* expose secrets for notification templates ([#3455](https://github.com/argoproj/argo-rollouts/issues/3455)) ([#3466](https://github.com/argoproj/argo-rollouts/issues/3466))
* ping pong support for istio ([#3371](https://github.com/argoproj/argo-rollouts/issues/3371))
* display init container images on the rollout dashboard ([#3473](https://github.com/argoproj/argo-rollouts/issues/3473))
* add Analysis run to rollout  notifications ([#3296](https://github.com/argoproj/argo-rollouts/issues/3296))
* add the max traffic weight support for the traffic routing (nginx/plugins). ([#3215](https://github.com/argoproj/argo-rollouts/issues/3215))
* allow analysis run to use separate kubeconfig for jobs ([#3350](https://github.com/argoproj/argo-rollouts/issues/3350))
* Support AnalysisRunMetadata and Dryrun for experiments via Rollout ([#3213](https://github.com/argoproj/argo-rollouts/issues/3213))
* allow setting traefik versions ([#3348](https://github.com/argoproj/argo-rollouts/issues/3348))
* support ability to run only the analysis controller ([#3336](https://github.com/argoproj/argo-rollouts/issues/3336))
* Reference AnalysisTemplates inside an AnalysisTemplate ([#3353](https://github.com/argoproj/argo-rollouts/issues/3353))
* Add support for aggregator type in DataDog metric provider  ([#3293](https://github.com/argoproj/argo-rollouts/issues/3293))
* release-1.6 logs for priceline
* release-1.6 logs for priceline
* release-1.6 logs for priceline
* release-1.6 logs for priceline
* add analysis modal ([#3174](https://github.com/argoproj/argo-rollouts/issues/3174))
* automatically scale down Deployment after migrating to Rollout ([#3111](https://github.com/argoproj/argo-rollouts/issues/3111))
* Rollouts UI List View Refresh ([#3118](https://github.com/argoproj/argo-rollouts/issues/3118))
* **analysis:** add ttlStrategy on AnalysisRun for garbage collecting stale AnalysisRun automatically ([#3324](https://github.com/argoproj/argo-rollouts/issues/3324))
* **dashboard:** improve pods visibility ([#3483](https://github.com/argoproj/argo-rollouts/issues/3483))
* **trafficrouting:** use values array for multiple accepted values under same header name ([#2974](https://github.com/argoproj/argo-rollouts/issues/2974))

### Fix

* docs site version selector broken ([#3590](https://github.com/argoproj/argo-rollouts/issues/3590))
* don't default datadog aggregator ([#3643](https://github.com/argoproj/argo-rollouts/issues/3643))
* Add volume for plugin and tmp folder ([#3546](https://github.com/argoproj/argo-rollouts/issues/3546))
* verify the weight of the alb at the end of the rollout ([#3627](https://github.com/argoproj/argo-rollouts/issues/3627))
* when Rollout has pingpong and stable/canary service defined, only alb traffic management uses pingpong. ([#3628](https://github.com/argoproj/argo-rollouts/issues/3628))
* protocol missing in ambassador canary mapping creation. Fixes  [#3593](https://github.com/argoproj/argo-rollouts/issues/3593) ([#3603](https://github.com/argoproj/argo-rollouts/issues/3603))
* rs conflict with fallback to patch ([#3559](https://github.com/argoproj/argo-rollouts/issues/3559))
* analysis step should be ignored after promote ([#3016](https://github.com/argoproj/argo-rollouts/issues/3016))
* set formatter for klog logger ([#3493](https://github.com/argoproj/argo-rollouts/issues/3493))
* fix the issue that when max weight is 100000000, and the replicas> 20,  the trafficWeightToReplicas will return negative value. ([#3474](https://github.com/argoproj/argo-rollouts/issues/3474))
* Add the GOPATH to the go-to-protobuf command ([#3022](https://github.com/argoproj/argo-rollouts/issues/3022))
* job metrics owner ref when using custom job kubeconfig/ns ([#3425](https://github.com/argoproj/argo-rollouts/issues/3425))
* prevent hot loop when fully promoted rollout is aborted ([#3064](https://github.com/argoproj/argo-rollouts/issues/3064))
* inopportune scaling events would lose some status fields ([#3060](https://github.com/argoproj/argo-rollouts/issues/3060))
* include the correct response error in the plugin init error message ([#3388](https://github.com/argoproj/argo-rollouts/issues/3388))
* append weighted destination only when weight is mentioned  ([#2734](https://github.com/argoproj/argo-rollouts/issues/2734))
* stuck rollout when 2nd deployment happens before 1st finishes ([#3354](https://github.com/argoproj/argo-rollouts/issues/3354))
* do not require pod readiness when switching desired service selector on abort ([#3338](https://github.com/argoproj/argo-rollouts/issues/3338))
* log rs name when update fails ([#3318](https://github.com/argoproj/argo-rollouts/issues/3318))
* keep rs inormer updated upon updating labels and annotations ([#3321](https://github.com/argoproj/argo-rollouts/issues/3321))
* canary step analysis run wasn't terminated as keep running after promote action being called. Fixes [#3220](https://github.com/argoproj/argo-rollouts/issues/3220) ([#3221](https://github.com/argoproj/argo-rollouts/issues/3221))
* updates to replicas and pod template at the same time causes rollout to get stuck ([#3272](https://github.com/argoproj/argo-rollouts/issues/3272))
* revert repo change to expr ([#3094](https://github.com/argoproj/argo-rollouts/issues/3094))
* make sure we use the updated rs when we write back to informer ([#3237](https://github.com/argoproj/argo-rollouts/issues/3237))
* conflict on updates to replicaset revision ([#3216](https://github.com/argoproj/argo-rollouts/issues/3216))
* rollouts getting stuck due to bad rs informer updates ([#3200](https://github.com/argoproj/argo-rollouts/issues/3200))
* Revert "fix: istio destionationrule subsets enforcement ([#3126](https://github.com/argoproj/argo-rollouts/issues/3126))" ([#3147](https://github.com/argoproj/argo-rollouts/issues/3147))
* istio destionationrule subsets enforcement ([#3126](https://github.com/argoproj/argo-rollouts/issues/3126))
* docs require build.os to be defined ([#3133](https://github.com/argoproj/argo-rollouts/issues/3133))
* rollback to stable with dynamicStableScale could overwhelm stable pods ([#3077](https://github.com/argoproj/argo-rollouts/issues/3077))
* sync notification controller configmaps/secrets first ([#3075](https://github.com/argoproj/argo-rollouts/issues/3075))
* codegen was missed ([#3104](https://github.com/argoproj/argo-rollouts/issues/3104))
* keep rs informer updated ([#3091](https://github.com/argoproj/argo-rollouts/issues/3091))
* missing notification on error ([#3076](https://github.com/argoproj/argo-rollouts/issues/3076))
* canary step analysis run wasn't terminated as keep running after promote action being called. Fixes [#3220](https://github.com/argoproj/argo-rollouts/issues/3220) ([#3221](https://github.com/argoproj/argo-rollouts/issues/3221))
* Replace antonmedv/expr with expr-lang/expr ([#3090](https://github.com/argoproj/argo-rollouts/issues/3090))
* bump notification-engine to fix double send on self server notifications ([#3095](https://github.com/argoproj/argo-rollouts/issues/3095))
* **controller:** typo fix ("Secrete" -> "Secret") in secret informer ([#2965](https://github.com/argoproj/argo-rollouts/issues/2965))
* **controller:** don't timeout rollout when still waiting for scale down delay ([#3417](https://github.com/argoproj/argo-rollouts/issues/3417))
* **controller:** treat spec.canary.analysis.template empty list as spec.canary.analysis not set ([#3446](https://github.com/argoproj/argo-rollouts/issues/3446))
* **controller:** prevent negative vsvc weights on a replica scaledown following a canary abort for istio trafficrouting ([#3467](https://github.com/argoproj/argo-rollouts/issues/3467))
* **controller:** Corrects the logic of comparing sha256 has. Fixes [#3519](https://github.com/argoproj/argo-rollouts/issues/3519) ([#3520](https://github.com/argoproj/argo-rollouts/issues/3520))
* **controller:** rollback should skip all steps to active rs within RollbackWindow ([#2953](https://github.com/argoproj/argo-rollouts/issues/2953))
* **metricprovider:** support Datadog v2 API Fixes [#2813](https://github.com/argoproj/argo-rollouts/issues/2813) ([#2997](https://github.com/argoproj/argo-rollouts/issues/2997))

### Refactor

* rename interface{} => any ([#3000](https://github.com/argoproj/argo-rollouts/issues/3000))

### Test

* add unit tests for maxSurge=0, replicas=1 ([#3375](https://github.com/argoproj/argo-rollouts/issues/3375))


<a name="1.6.1"></a>
## [1.6.1](https://github.com/argoproj/argo-rollouts/compare/v1.6.1-cap-CR-update-to1.6.1...1.6.1) (2023-12-07)


<a name="v1.6.1-cap-CR-update-to1.6.1"></a>
## [v1.6.1-cap-CR-update-to1.6.1](https://github.com/argoproj/argo-rollouts/compare/1.6.1-analysis-fix...v1.6.1-cap-CR-update-to1.6.1) (2023-12-05)


<a name="1.6.1-analysis-fix"></a>
## [1.6.1-analysis-fix](https://github.com/argoproj/argo-rollouts/compare/1.6.1-additional-logs...1.6.1-analysis-fix) (2024-01-18)


<a name="1.6.1-additional-logs"></a>
## [1.6.1-additional-logs](https://github.com/argoproj/argo-rollouts/compare/v1.6.1-CR-23199...1.6.1-additional-logs) (2024-02-06)


<a name="v1.6.1-CR-23199"></a>
## [v1.6.1-CR-23199](https://github.com/argoproj/argo-rollouts/compare/1.5.0...v1.6.1-CR-23199) (2024-07-01)

### Build

* **deps:** always resolve momentjs version 2.29.4 ([#3182](https://github.com/argoproj/argo-rollouts/issues/3182))

### Chore

* quote golang version string to not use go 1.2.2 ([#2915](https://github.com/argoproj/argo-rollouts/issues/2915))
* upgrade cosign ([#3139](https://github.com/argoproj/argo-rollouts/issues/3139))
* add missing rollout fields ([#3062](https://github.com/argoproj/argo-rollouts/issues/3062))
* change file name for readthedocs compatibility ([#2999](https://github.com/argoproj/argo-rollouts/issues/2999))
* Update test and related docs for plugin name standard ([#2728](https://github.com/argoproj/argo-rollouts/issues/2728))
* bump k8s deps to v0.25.8 ([#2712](https://github.com/argoproj/argo-rollouts/issues/2712))
* add zachaller as lead in owers file ([#2759](https://github.com/argoproj/argo-rollouts/issues/2759))
* bump gotestsum and fix flakey test causing nil channel send ([#2934](https://github.com/argoproj/argo-rollouts/issues/2934))
* Add tests for pause functionality in rollout package ([#2772](https://github.com/argoproj/argo-rollouts/issues/2772))
* add unit test ([#2798](https://github.com/argoproj/argo-rollouts/issues/2798))
* leave the validation of setHeaderRoute to the plugin when plugins is not empty. ([#2898](https://github.com/argoproj/argo-rollouts/issues/2898))
* bump golang to 1.20 ([#2910](https://github.com/argoproj/argo-rollouts/issues/2910))
* add make help cmd ([#2854](https://github.com/argoproj/argo-rollouts/issues/2854))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.26.1 to 1.26.2 ([#2848](https://github.com/argoproj/argo-rollouts/issues/2848))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.31 to 1.18.32 ([#2928](https://github.com/argoproj/argo-rollouts/issues/2928))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.26.3 to 1.27.0 ([#2922](https://github.com/argoproj/argo-rollouts/issues/2922))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.30 to 1.18.31 ([#2924](https://github.com/argoproj/argo-rollouts/issues/2924))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.29 to 1.18.30 ([#2919](https://github.com/argoproj/argo-rollouts/issues/2919))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.19.0 to 1.19.1 ([#2920](https://github.com/argoproj/argo-rollouts/issues/2920))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.0 to 1.27.1 ([#2927](https://github.com/argoproj/argo-rollouts/issues/2927))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.20.1 to 1.20.2 ([#2941](https://github.com/argoproj/argo-rollouts/issues/2941))
* **deps:** bump google.golang.org/grpc from 1.56.2 to 1.57.0 ([#2908](https://github.com/argoproj/argo-rollouts/issues/2908))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.28 to 1.18.29 ([#2907](https://github.com/argoproj/argo-rollouts/issues/2907))
* **deps:** bump github.com/antonmedv/expr from 1.12.6 to 1.12.7 ([#2894](https://github.com/argoproj/argo-rollouts/issues/2894))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.26.2 to 1.26.3 ([#2884](https://github.com/argoproj/argo-rollouts/issues/2884))
* **deps:** bump docker/setup-qemu-action from 2.1.0 to 2.2.0 ([#2878](https://github.com/argoproj/argo-rollouts/issues/2878))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.27 to 1.18.28 ([#2883](https://github.com/argoproj/argo-rollouts/issues/2883))
* **deps:** bump slsa-framework/slsa-github-generator from 1.6.0 to 1.7.0 ([#2880](https://github.com/argoproj/argo-rollouts/issues/2880))
* **deps:** bump actions/setup-go from 4.0.0 to 4.0.1 ([#2881](https://github.com/argoproj/argo-rollouts/issues/2881))
* **deps:** bump docker/setup-buildx-action from 2.5.0 to 2.9.1 ([#2879](https://github.com/argoproj/argo-rollouts/issues/2879))
* **deps:** bump docker/login-action from 2.1.0 to 2.2.0 ([#2877](https://github.com/argoproj/argo-rollouts/issues/2877))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.13 to 1.19.14 ([#2886](https://github.com/argoproj/argo-rollouts/issues/2886))
* **deps:** bump github.com/antonmedv/expr from 1.12.5 to 1.12.6 ([#2882](https://github.com/argoproj/argo-rollouts/issues/2882))
* **deps:** bump google.golang.org/grpc from 1.56.1 to 1.56.2 ([#2872](https://github.com/argoproj/argo-rollouts/issues/2872))
* **deps:** bump sigstore/cosign-installer from 3.1.0 to 3.1.1 ([#2860](https://github.com/argoproj/argo-rollouts/issues/2860))
* **deps:** bump google.golang.org/protobuf from 1.30.0 to 1.31.0 ([#2859](https://github.com/argoproj/argo-rollouts/issues/2859))
* **deps:** bump sigstore/cosign-installer from 3.0.5 to 3.1.0 ([#2858](https://github.com/argoproj/argo-rollouts/issues/2858))
* **deps:** bump google.golang.org/grpc from 1.55.0 to 1.56.1 ([#2856](https://github.com/argoproj/argo-rollouts/issues/2856))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.14 to 1.20.1 ([#2926](https://github.com/argoproj/argo-rollouts/issues/2926))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.12 to 1.19.13 ([#2847](https://github.com/argoproj/argo-rollouts/issues/2847))
* **deps:** bump actions/setup-go from 3.5.0 to 4.0.1 ([#2849](https://github.com/argoproj/argo-rollouts/issues/2849))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.26 to 1.18.27 ([#2844](https://github.com/argoproj/argo-rollouts/issues/2844))
* **deps:** bump github.com/prometheus/client_golang from 1.15.1 to 1.16.0 ([#2846](https://github.com/argoproj/argo-rollouts/issues/2846))
* **deps:** bump slsa-framework/slsa-github-generator from 1.7.0 to 1.8.0 ([#2936](https://github.com/argoproj/argo-rollouts/issues/2936))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.11 to 1.19.12 ([#2839](https://github.com/argoproj/argo-rollouts/issues/2839))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.26.0 to 1.26.1 ([#2840](https://github.com/argoproj/argo-rollouts/issues/2840))
* **deps:** bump docker/build-push-action from 4.1.0 to 4.1.1 ([#2837](https://github.com/argoproj/argo-rollouts/issues/2837))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.25 to 1.18.26 ([#2841](https://github.com/argoproj/argo-rollouts/issues/2841))
* **deps:** bump docker/build-push-action from 4.0.0 to 4.1.0 ([#2832](https://github.com/argoproj/argo-rollouts/issues/2832))
* **deps:** bump github.com/sirupsen/logrus from 1.9.2 to 1.9.3 ([#2821](https://github.com/argoproj/argo-rollouts/issues/2821))
* **deps:** bump github.com/hashicorp/go-plugin from 1.4.9 to 1.4.10 ([#2822](https://github.com/argoproj/argo-rollouts/issues/2822))
* **deps:** bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#2817](https://github.com/argoproj/argo-rollouts/issues/2817))
* **deps:** bump github.com/sirupsen/logrus from 1.9.1 to 1.9.2 ([#2789](https://github.com/argoproj/argo-rollouts/issues/2789))
* **deps:** bump github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#2796](https://github.com/argoproj/argo-rollouts/issues/2796))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.27.1 to 1.27.2 ([#2944](https://github.com/argoproj/argo-rollouts/issues/2944))
* **deps:** bump sigstore/cosign-installer from 3.0.3 to 3.0.5 ([#2788](https://github.com/argoproj/argo-rollouts/issues/2788))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.32 to 1.18.33 ([#2943](https://github.com/argoproj/argo-rollouts/issues/2943))
* **deps:** bump github.com/sirupsen/logrus from 1.9.0 to 1.9.1 ([#2784](https://github.com/argoproj/argo-rollouts/issues/2784))
* **deps:** bump codecov/codecov-action from 3.1.3 to 3.1.4 ([#2782](https://github.com/argoproj/argo-rollouts/issues/2782))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.24 to 1.18.25 ([#2770](https://github.com/argoproj/argo-rollouts/issues/2770))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.23 to 1.18.24 ([#2768](https://github.com/argoproj/argo-rollouts/issues/2768))
* **deps:** bump google.golang.org/grpc from 1.54.0 to 1.55.0 ([#2763](https://github.com/argoproj/argo-rollouts/issues/2763))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.22 to 1.18.23 ([#2756](https://github.com/argoproj/argo-rollouts/issues/2756))
* **deps:** bump actions/setup-go from 4.0.1 to 4.1.0 ([#2947](https://github.com/argoproj/argo-rollouts/issues/2947))
* **deps:** replace `github.com/ghodss/yaml` with `sigs.k8s.io/yaml` ([#2681](https://github.com/argoproj/argo-rollouts/issues/2681))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.10 to 1.26.0 ([#2755](https://github.com/argoproj/argo-rollouts/issues/2755))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.10 to 1.19.11 ([#2757](https://github.com/argoproj/argo-rollouts/issues/2757))
* **deps:** bump github.com/prometheus/client_golang from 1.15.0 to 1.15.1 ([#2754](https://github.com/argoproj/argo-rollouts/issues/2754))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.21 to 1.18.22 ([#2746](https://github.com/argoproj/argo-rollouts/issues/2746))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.9 to 1.25.10 ([#2745](https://github.com/argoproj/argo-rollouts/issues/2745))
* **deps:** bump github.com/antonmedv/expr from 1.12.7 to 1.13.0 ([#2951](https://github.com/argoproj/argo-rollouts/issues/2951))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.9 to 1.19.10 ([#2747](https://github.com/argoproj/argo-rollouts/issues/2747))
* **deps:** bump codecov/codecov-action from 3.1.2 to 3.1.3 ([#2735](https://github.com/argoproj/argo-rollouts/issues/2735))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.20.2 to 1.21.0 ([#2950](https://github.com/argoproj/argo-rollouts/issues/2950))
* **deps:** bump github.com/prometheus/client_golang from 1.14.0 to 1.15.0 ([#2721](https://github.com/argoproj/argo-rollouts/issues/2721))
* **deps:** bump codecov/codecov-action from 3.1.1 to 3.1.2 ([#2711](https://github.com/argoproj/argo-rollouts/issues/2711))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.20 to 1.18.21 ([#2709](https://github.com/argoproj/argo-rollouts/issues/2709))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.8 to 1.19.9 ([#2708](https://github.com/argoproj/argo-rollouts/issues/2708))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.8 to 1.25.9 ([#2710](https://github.com/argoproj/argo-rollouts/issues/2710))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.19 to 1.18.20 ([#2705](https://github.com/argoproj/argo-rollouts/issues/2705))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.7 to 1.19.8 ([#2704](https://github.com/argoproj/argo-rollouts/issues/2704))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.17.7 to 1.17.8 ([#2703](https://github.com/argoproj/argo-rollouts/issues/2703))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.7 to 1.25.8 ([#2702](https://github.com/argoproj/argo-rollouts/issues/2702))
* **deps:** bump peter-evans/create-pull-request from 4 to 5 ([#2697](https://github.com/argoproj/argo-rollouts/issues/2697))
* **deps:** bump github.com/spf13/cobra from 1.6.1 to 1.7.0 ([#2698](https://github.com/argoproj/argo-rollouts/issues/2698))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 from 2.12.2 to 2.12.3 ([#2684](https://github.com/argoproj/argo-rollouts/issues/2684))

### Ci

* generate attestations during a release  ([#2785](https://github.com/argoproj/argo-rollouts/issues/2785))
* use keyless signing for main and release branches ([#2783](https://github.com/argoproj/argo-rollouts/issues/2783))

### Docs

* update supported k8s version ([#2949](https://github.com/argoproj/argo-rollouts/issues/2949))
* mirroring support in Traefik is not implemented yet ([#2904](https://github.com/argoproj/argo-rollouts/issues/2904))
* update contributions.md to include k3d as recommended cluster, add details on e2e test setup, and update kubectl install link. Fixes [#1750](https://github.com/argoproj/argo-rollouts/issues/1750) ([#1867](https://github.com/argoproj/argo-rollouts/issues/1867))
* fix minor mistakes in Migrating to Deployments ([#2270](https://github.com/argoproj/argo-rollouts/issues/2270))
* Update docs of Rollout spec to add active/previewMetadata ([#2833](https://github.com/argoproj/argo-rollouts/issues/2833))
* Update datadog.md - clarify formulas [#2813](https://github.com/argoproj/argo-rollouts/issues/2813) ([#2819](https://github.com/argoproj/argo-rollouts/issues/2819))
* support for Kong ingress ([#2820](https://github.com/argoproj/argo-rollouts/issues/2820))
* Fix AWS App Mesh getting started documentation to avoid connection pooling problems ([#2814](https://github.com/argoproj/argo-rollouts/issues/2814))
* Update Changelog ([#2807](https://github.com/argoproj/argo-rollouts/issues/2807))
* use correct capitalization for "Datadog" in navigation sidebar ([#2809](https://github.com/argoproj/argo-rollouts/issues/2809))
* Fix typo in header routing specification docs ([#2808](https://github.com/argoproj/argo-rollouts/issues/2808))
* support for Google Cloud Load balancers ([#2803](https://github.com/argoproj/argo-rollouts/issues/2803))
* Show how plugins are loaded ([#2801](https://github.com/argoproj/argo-rollouts/issues/2801))
* Add gateway API link, fix Contour plugin naming ([#2787](https://github.com/argoproj/argo-rollouts/issues/2787))
* Add some details around running locally to make things clearer new contributors ([#2786](https://github.com/argoproj/argo-rollouts/issues/2786))
* Add docs for Amazon Managed Prometheus ([#2777](https://github.com/argoproj/argo-rollouts/issues/2777))
* Update Changelog ([#2765](https://github.com/argoproj/argo-rollouts/issues/2765))
* copy argo cd docs drop down fix ([#2731](https://github.com/argoproj/argo-rollouts/issues/2731))
* Add contour trafficrouter plugin ([#2729](https://github.com/argoproj/argo-rollouts/issues/2729))
* fix link to plugins for traffic routers ([#2719](https://github.com/argoproj/argo-rollouts/issues/2719))
* Update Changelog ([#2683](https://github.com/argoproj/argo-rollouts/issues/2683))
* **analysis:** fix use stringData in the examples ([#2715](https://github.com/argoproj/argo-rollouts/issues/2715))
* **example:** Add example on how to execute subset of e2e tests ([#2867](https://github.com/argoproj/argo-rollouts/issues/2867))
* **example:** interval requires count ([#2690](https://github.com/argoproj/argo-rollouts/issues/2690))

### Feat

* release-1.6 logs for priceline
* release-1.6 logs for priceline
* release-1.6 logs for priceline
* release-1.6 logs for priceline
* release-1.6
* release-1.6
* release-1.6
* retain TLS configuration for canary ingresses in the nginx integration. Fixes [#1134](https://github.com/argoproj/argo-rollouts/issues/1134) ([#2679](https://github.com/argoproj/argo-rollouts/issues/2679))
* enable self service notification support ([#2930](https://github.com/argoproj/argo-rollouts/issues/2930))
* support prometheus headers ([#2937](https://github.com/argoproj/argo-rollouts/issues/2937))
* Add insecure option for Prometheus. Fixes [#2913](https://github.com/argoproj/argo-rollouts/issues/2913) ([#2914](https://github.com/argoproj/argo-rollouts/issues/2914))
* Add prometheus timeout ([#2893](https://github.com/argoproj/argo-rollouts/issues/2893))
* Support Multiple ALB Ingresses ([#2639](https://github.com/argoproj/argo-rollouts/issues/2639))
* Send informer add k8s event ([#2834](https://github.com/argoproj/argo-rollouts/issues/2834))
* add merge key to analysis template ([#2842](https://github.com/argoproj/argo-rollouts/issues/2842))
* **analysis:** Adds rollout Spec.Selector.MatchLabels to AnalysisRun. Fixes [#2888](https://github.com/argoproj/argo-rollouts/issues/2888) ([#2903](https://github.com/argoproj/argo-rollouts/issues/2903))
* **controller:** Add custom metadata support for AnalysisRun. Fixes [#2740](https://github.com/argoproj/argo-rollouts/issues/2740) ([#2743](https://github.com/argoproj/argo-rollouts/issues/2743))
* **dashboard:** Refresh Rollouts dashboard UI ([#2723](https://github.com/argoproj/argo-rollouts/issues/2723))
* **metricprovider:** allow user to define metrics.provider.job.metadata ([#2762](https://github.com/argoproj/argo-rollouts/issues/2762))

### Fix

* resolve args to metric in garbage collection function ([#2843](https://github.com/argoproj/argo-rollouts/issues/2843))
* rollouts getting stuck due to bad rs informer updates ([#3200](https://github.com/argoproj/argo-rollouts/issues/3200))
* Revert "fix: istio destionationrule subsets enforcement ([#3126](https://github.com/argoproj/argo-rollouts/issues/3126))" ([#3147](https://github.com/argoproj/argo-rollouts/issues/3147))
* istio destionationrule subsets enforcement ([#3126](https://github.com/argoproj/argo-rollouts/issues/3126))
* docs require build.os to be defined ([#3133](https://github.com/argoproj/argo-rollouts/issues/3133))
* inopportune scaling events would lose some status fields ([#3060](https://github.com/argoproj/argo-rollouts/issues/3060))
* rollback to stable with dynamicStableScale could overwhelm stable pods ([#3077](https://github.com/argoproj/argo-rollouts/issues/3077))
* prevent hot loop when fully promoted rollout is aborted ([#3064](https://github.com/argoproj/argo-rollouts/issues/3064))
* keep rs informer updated ([#3091](https://github.com/argoproj/argo-rollouts/issues/3091))
* bump notification-engine to fix double send on self server notifications ([#3095](https://github.com/argoproj/argo-rollouts/issues/3095))
* sync notification controller configmaps/secrets first ([#3075](https://github.com/argoproj/argo-rollouts/issues/3075))
* missing notification on error ([#3076](https://github.com/argoproj/argo-rollouts/issues/3076))
* analysis step should be ignored after promote ([#3016](https://github.com/argoproj/argo-rollouts/issues/3016))
* change logic of analysis run to better handle errors ([#2695](https://github.com/argoproj/argo-rollouts/issues/2695))
* istio dropping fields during removing of managed routes ([#2692](https://github.com/argoproj/argo-rollouts/issues/2692))
* add required ingress permission ([#2933](https://github.com/argoproj/argo-rollouts/issues/2933))
* cloudwatch metrics provider multiple dimensions ([#2932](https://github.com/argoproj/argo-rollouts/issues/2932))
*  rollout not modify the VirtualService whit setHeaderRoute step with workloadRef ([#2797](https://github.com/argoproj/argo-rollouts/issues/2797))
* get new httpRoutesI after removeRoute() to avoid duplicates. Fixes [#2769](https://github.com/argoproj/argo-rollouts/issues/2769) ([#2887](https://github.com/argoproj/argo-rollouts/issues/2887))
* properly wrap Datadog API v2 request body ([#2771](https://github.com/argoproj/argo-rollouts/issues/2771)) ([#2775](https://github.com/argoproj/argo-rollouts/issues/2775))
* make new alb fullName field  optional for backward compatability ([#2806](https://github.com/argoproj/argo-rollouts/issues/2806))
* canary step analysis run wasn't terminated as keep running after promote action being called. Fixes [#3220](https://github.com/argoproj/argo-rollouts/issues/3220) ([#3221](https://github.com/argoproj/argo-rollouts/issues/3221))
* **analysis:** Adding field in YAML to provide region for Sigv4 signing.  ([#2794](https://github.com/argoproj/argo-rollouts/issues/2794))
* **analysis:** Graphite query - remove whitespaces ([#2752](https://github.com/argoproj/argo-rollouts/issues/2752))
* **analysis:** Graphite metric provider - index out of range [0] with length 0 ([#2751](https://github.com/argoproj/argo-rollouts/issues/2751))
* **controller:** Remove name label from some k8s client metrics on events and replicasets ([#2851](https://github.com/argoproj/argo-rollouts/issues/2851))
* **controller:** Add klog logrus bridge. Fixes [#2707](https://github.com/argoproj/argo-rollouts/issues/2707). ([#2701](https://github.com/argoproj/argo-rollouts/issues/2701))
* **controller:** typo fix ("Secrete" -> "Secret") in secret informer ([#2965](https://github.com/argoproj/argo-rollouts/issues/2965))
* **controller:** Fix for rollouts getting stuck in loop ([#2689](https://github.com/argoproj/argo-rollouts/issues/2689))
* **controller:** rollback should skip all steps to active rs within RollbackWindow ([#2953](https://github.com/argoproj/argo-rollouts/issues/2953))
* **trafficrouting:** apply stable selectors on canary service on rollout abort [#2781](https://github.com/argoproj/argo-rollouts/issues/2781) ([#2818](https://github.com/argoproj/argo-rollouts/issues/2818))

### Refactor

* change plugin naming pattern [#2720](https://github.com/argoproj/argo-rollouts/issues/2720) ([#2722](https://github.com/argoproj/argo-rollouts/issues/2722))

### BREAKING CHANGE


The metric labels have changed on controller_clientset_k8s_request_total to not include the name of the resource for events and replicasets. These names have generated hashes in them and cause really high cardinality.

Remove name label from k8s some client metrics

The `name` label in the `controller_clientset_k8s_request_total` metric
produce an excessive amount of cardinality for `events` and `replicasets`.
This can lead to hundreds of thousands of unique metrics over a couple
weeks in a large deployment. Set the name to "N/A" for these client request
types.


<a name="1.5.0"></a>
## [1.5.0](https://github.com/argoproj/argo-rollouts/compare/v1.5.0...1.5.0) (2023-07-19)


<a name="v1.5.0"></a>
## [v1.5.0](https://github.com/argoproj/argo-rollouts/compare/v1.4.0-cap-sw...v1.5.0) (2023-07-19)

### Build

* manually run auto changelog and fix workflow ([#2494](https://github.com/argoproj/argo-rollouts/issues/2494))

### Chore

* update e2e k8s versions ([#2637](https://github.com/argoproj/argo-rollouts/issues/2637))
* Remove namespaced crds ([#2516](https://github.com/argoproj/argo-rollouts/issues/2516))
* fix dependabot broken dependency ([#2529](https://github.com/argoproj/argo-rollouts/issues/2529))
* disable docker sbom and attestations ([#2528](https://github.com/argoproj/argo-rollouts/issues/2528))
* improve e2e test timing ([#2577](https://github.com/argoproj/argo-rollouts/issues/2577))
* fix typo for json tag on rollbackWindow ([#2598](https://github.com/argoproj/argo-rollouts/issues/2598))
* update package dependencie ([#2602](https://github.com/argoproj/argo-rollouts/issues/2602))
* bump node version and set openssl-legacy-provider ([#2606](https://github.com/argoproj/argo-rollouts/issues/2606))
* bump k8s deps to v0.25.8 ([#2712](https://github.com/argoproj/argo-rollouts/issues/2712))
* switch to distroless for cli/dashboard image ([#2596](https://github.com/argoproj/argo-rollouts/issues/2596))
* add Tuhu to users ([#2630](https://github.com/argoproj/argo-rollouts/issues/2630))
* bump deps for prisma ([#2643](https://github.com/argoproj/argo-rollouts/issues/2643))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.6 to 1.19.7 ([#2672](https://github.com/argoproj/argo-rollouts/issues/2672))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.15 to 1.18.16 ([#2652](https://github.com/argoproj/argo-rollouts/issues/2652))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.16 to 1.18.17 ([#2659](https://github.com/argoproj/argo-rollouts/issues/2659))
* **deps:** bump github.com/antonmedv/expr from 1.12.2 to 1.12.3 ([#2653](https://github.com/argoproj/argo-rollouts/issues/2653))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.5 to 1.19.6 ([#2654](https://github.com/argoproj/argo-rollouts/issues/2654))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.4 to 1.25.5 ([#2655](https://github.com/argoproj/argo-rollouts/issues/2655))
* **deps:** bump github.com/antonmedv/expr from 1.12.1 to 1.12.2 ([#2649](https://github.com/argoproj/argo-rollouts/issues/2649))
* **deps:** bump google.golang.org/protobuf from 1.28.1 to 1.29.0 ([#2646](https://github.com/argoproj/argo-rollouts/issues/2646))
* **deps:** bump github.com/golang/protobuf from 1.5.2 to 1.5.3 ([#2645](https://github.com/argoproj/argo-rollouts/issues/2645))
* **deps:** bump github.com/prometheus/common from 0.41.0 to 0.42.0 ([#2644](https://github.com/argoproj/argo-rollouts/issues/2644))
* **deps:** bump minimist from 1.2.5 to 1.2.8 in /ui ([#2638](https://github.com/argoproj/argo-rollouts/issues/2638))
* **deps:** bump github.com/hashicorp/go-plugin from 1.4.8 to 1.4.9 ([#2636](https://github.com/argoproj/argo-rollouts/issues/2636))
* **deps:** bump github.com/prometheus/common from 0.40.0 to 0.41.0 ([#2634](https://github.com/argoproj/argo-rollouts/issues/2634))
* **deps:** bump google.golang.org/protobuf from 1.29.0 to 1.29.1 ([#2660](https://github.com/argoproj/argo-rollouts/issues/2660))
* **deps:** bump google.golang.org/protobuf from 1.29.1 to 1.30.0 ([#2665](https://github.com/argoproj/argo-rollouts/issues/2665))
* **deps:** bump github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#2627](https://github.com/argoproj/argo-rollouts/issues/2627))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.14 to 1.18.15 ([#2618](https://github.com/argoproj/argo-rollouts/issues/2618))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.3 to 1.25.4 ([#2617](https://github.com/argoproj/argo-rollouts/issues/2617))
* **deps:** bump github.com/antonmedv/expr from 1.12.0 to 1.12.1 ([#2619](https://github.com/argoproj/argo-rollouts/issues/2619))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.4 to 1.19.5 ([#2616](https://github.com/argoproj/argo-rollouts/issues/2616))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.19.3 to 1.19.4 ([#2612](https://github.com/argoproj/argo-rollouts/issues/2612))
* **deps:** bump github.com/prometheus/common from 0.39.0 to 0.40.0 ([#2611](https://github.com/argoproj/argo-rollouts/issues/2611))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.13 to 1.18.14 ([#2614](https://github.com/argoproj/argo-rollouts/issues/2614))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.2 to 1.25.3 ([#2615](https://github.com/argoproj/argo-rollouts/issues/2615))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config from 1.18.17 to 1.18.19 ([#2673](https://github.com/argoproj/argo-rollouts/issues/2673))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.5 to 1.25.6 ([#2671](https://github.com/argoproj/argo-rollouts/issues/2671))
* **deps:** bump imjasonh/setup-crane from 0.2 to 0.3 ([#2600](https://github.com/argoproj/argo-rollouts/issues/2600))
* **deps:** bump github.com/antonmedv/expr from 1.12.3 to 1.12.5 ([#2670](https://github.com/argoproj/argo-rollouts/issues/2670))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2593](https://github.com/argoproj/argo-rollouts/issues/2593))
* **deps:** bump google.golang.org/grpc from 1.53.0 to 1.54.0 ([#2674](https://github.com/argoproj/argo-rollouts/issues/2674))
* **deps:** bump google.golang.org/grpc from 1.52.3 to 1.53.0 ([#2574](https://github.com/argoproj/argo-rollouts/issues/2574))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2565](https://github.com/argoproj/argo-rollouts/issues/2565))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2564](https://github.com/argoproj/argo-rollouts/issues/2564))
* **deps:** bump github.com/antonmedv/expr from 1.11.0 to 1.12.0 ([#2567](https://github.com/argoproj/argo-rollouts/issues/2567))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.17.3 to 1.17.4 ([#2566](https://github.com/argoproj/argo-rollouts/issues/2566))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2563](https://github.com/argoproj/argo-rollouts/issues/2563))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2559](https://github.com/argoproj/argo-rollouts/issues/2559))
* **deps:** bump github.com/antonmedv/expr from 1.9.0 to 1.11.0 ([#2558](https://github.com/argoproj/argo-rollouts/issues/2558))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2555](https://github.com/argoproj/argo-rollouts/issues/2555))
* **deps:** bump docker/build-push-action from 3.3.0 to 4.0.0 ([#2550](https://github.com/argoproj/argo-rollouts/issues/2550))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 ([#2544](https://github.com/argoproj/argo-rollouts/issues/2544))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2542](https://github.com/argoproj/argo-rollouts/issues/2542))
* **deps:** bump google.golang.org/grpc from 1.52.1 to 1.52.3 ([#2541](https://github.com/argoproj/argo-rollouts/issues/2541))
* **deps:** bump google.golang.org/grpc from 1.52.0 to 1.52.1 ([#2538](https://github.com/argoproj/argo-rollouts/issues/2538))
* **deps:** bump dependabot/fetch-metadata from 1.3.5 to 1.3.6 ([#2537](https://github.com/argoproj/argo-rollouts/issues/2537))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2534](https://github.com/argoproj/argo-rollouts/issues/2534))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2533](https://github.com/argoproj/argo-rollouts/issues/2533))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2532](https://github.com/argoproj/argo-rollouts/issues/2532))
* **deps:** bump actions/setup-go from 3 to 4 ([#2663](https://github.com/argoproj/argo-rollouts/issues/2663))
* **deps:** bump actions/stale from 7 to 8 ([#2677](https://github.com/argoproj/argo-rollouts/issues/2677))
* **deps:** bump github.com/antonmedv/expr from 1.9.0 to 1.10.0 ([#2527](https://github.com/argoproj/argo-rollouts/issues/2527))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2523](https://github.com/argoproj/argo-rollouts/issues/2523))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch from 1.25.6 to 1.25.7 ([#2682](https://github.com/argoproj/argo-rollouts/issues/2682))
* **deps:** bump google.golang.org/grpc from 1.51.0 to 1.52.0 ([#2513](https://github.com/argoproj/argo-rollouts/issues/2513))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2505](https://github.com/argoproj/argo-rollouts/issues/2505))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2506](https://github.com/argoproj/argo-rollouts/issues/2506))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2504](https://github.com/argoproj/argo-rollouts/issues/2504))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2497](https://github.com/argoproj/argo-rollouts/issues/2497))
* **deps:** bump actions/stale from 6 to 7 ([#2496](https://github.com/argoproj/argo-rollouts/issues/2496))

### Ci

* generate attestations during a release  ([#2785](https://github.com/argoproj/argo-rollouts/issues/2785))
* use keyless signing for main and release branches ([#2783](https://github.com/argoproj/argo-rollouts/issues/2783))

### Docs

* fix link to plugins for traffic routers ([#2719](https://github.com/argoproj/argo-rollouts/issues/2719))
* copy argo cd docs drop down fix ([#2731](https://github.com/argoproj/argo-rollouts/issues/2731))
* Mention Internet Bug Bounty in the security policy ([#2642](https://github.com/argoproj/argo-rollouts/issues/2642))
* Update Changelog ([#2625](https://github.com/argoproj/argo-rollouts/issues/2625))
* fix missing links for getting started documentation ([#2557](https://github.com/argoproj/argo-rollouts/issues/2557))
* fix spelling in example notification templates ([#2554](https://github.com/argoproj/argo-rollouts/issues/2554))
* Add best practice for reducing memory usage ([#2545](https://github.com/argoproj/argo-rollouts/issues/2545))
* commit generated docs for readthedocs.org ([#2535](https://github.com/argoproj/argo-rollouts/issues/2535))
* fix incorrect description for autoPromotionSeconds ([#2525](https://github.com/argoproj/argo-rollouts/issues/2525))
* manually add changelog due to action failure ([#2510](https://github.com/argoproj/argo-rollouts/issues/2510))
* fix typo apisix ([#2508](https://github.com/argoproj/argo-rollouts/issues/2508))
* add release schedule ([#2446](https://github.com/argoproj/argo-rollouts/issues/2446))
* fix rendering by upgrading deps ([#2495](https://github.com/argoproj/argo-rollouts/issues/2495))

### Feat

* Apache APISIX SetHeader support. Fixes [#2668](https://github.com/argoproj/argo-rollouts/issues/2668) ([#2678](https://github.com/argoproj/argo-rollouts/issues/2678))
* support N nginx ingresses ([#2467](https://github.com/argoproj/argo-rollouts/issues/2467))
* Add Service field to Rollout Experiment to allow service creation ([#2633](https://github.com/argoproj/argo-rollouts/issues/2633))
* Provide time.Parse and time.Now while evaluating notification trigger condition ([#2206](https://github.com/argoproj/argo-rollouts/issues/2206))
* Allow switching between Datadog v1 and v2. Fixes [#2549](https://github.com/argoproj/argo-rollouts/issues/2549) ([#2592](https://github.com/argoproj/argo-rollouts/issues/2592))
* add support for traffic router plugins ([#2573](https://github.com/argoproj/argo-rollouts/issues/2573))
* Add name attribute to ServicePort ([#2572](https://github.com/argoproj/argo-rollouts/issues/2572))
* metric plugin system based on hashicorp go-plugin ([#2514](https://github.com/argoproj/argo-rollouts/issues/2514))
* Adding SigV4 Option for Prometheus Metric Analysis ([#2489](https://github.com/argoproj/argo-rollouts/issues/2489))
* **analysis:** add Apache SkyWalking as metrics provider
* **controller:** Adding status.alb.canaryTargetGroup.fullName for ALB. Fixes [#2589](https://github.com/argoproj/argo-rollouts/issues/2589) ([#2604](https://github.com/argoproj/argo-rollouts/issues/2604))

### Fix

* resolve args to metric in garbage collection function ([#2843](https://github.com/argoproj/argo-rollouts/issues/2843))
* make new alb fullName field  optional for backward compatability ([#2806](https://github.com/argoproj/argo-rollouts/issues/2806))
* properly wrap Datadog API v2 request body ([#2771](https://github.com/argoproj/argo-rollouts/issues/2771)) ([#2775](https://github.com/argoproj/argo-rollouts/issues/2775))
* istio dropping fields during removing of managed routes ([#2692](https://github.com/argoproj/argo-rollouts/issues/2692))
* analysis information box [#2530](https://github.com/argoproj/argo-rollouts/issues/2530)  ([#2575](https://github.com/argoproj/argo-rollouts/issues/2575))
* change logic of analysis run to better handle errors ([#2695](https://github.com/argoproj/argo-rollouts/issues/2695))
* update GetTargetGroupMetadata to call DescribeTags in batches ([#2570](https://github.com/argoproj/argo-rollouts/issues/2570))
* remove outdated ioutil package dependencies ([#2583](https://github.com/argoproj/argo-rollouts/issues/2583))
* switch service selector back to stable on canary service when aborted ([#2540](https://github.com/argoproj/argo-rollouts/issues/2540))
* change log generator to only add CHANGELOG.md ([#2626](https://github.com/argoproj/argo-rollouts/issues/2626))
* Rollback change on service creation with weightless experiments ([#2624](https://github.com/argoproj/argo-rollouts/issues/2624))
* flakey TestWriteBackToInformer test ([#2621](https://github.com/argoproj/argo-rollouts/issues/2621))
* support only tls in virtual services ([#2502](https://github.com/argoproj/argo-rollouts/issues/2502))
* **analysis:** Nil Pointer Fixes [#2458](https://github.com/argoproj/argo-rollouts/issues/2458) ([#2680](https://github.com/argoproj/argo-rollouts/issues/2680))
* **controller:** Add klog logrus bridge. Fixes [#2707](https://github.com/argoproj/argo-rollouts/issues/2707). ([#2701](https://github.com/argoproj/argo-rollouts/issues/2701))
* **controller:** Fix for rollouts getting stuck in loop ([#2689](https://github.com/argoproj/argo-rollouts/issues/2689))

### BREAKING CHANGE


There was an unintentional change in behavior related to service creation with experiments introduced in v1.4.0 this has been reverted in v1.4.1 back to the original behavior. In v1.4.0 services where always created with for inline experiments even if there was no weight set. In 1.4.1 we go back to the original behavior of requiring weight to be set in order to create a service.


<a name="v1.4.0-cap-sw"></a>
## [v1.4.0-cap-sw](https://github.com/argoproj/argo-rollouts/compare/v1.4.0-cap-CR-10626...v1.4.0-cap-sw) (2023-01-10)

### Chore

* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2492](https://github.com/argoproj/argo-rollouts/issues/2492))


<a name="v1.4.0-cap-CR-10626"></a>
## [v1.4.0-cap-CR-10626](https://github.com/argoproj/argo-rollouts/compare/v1.2.0...v1.4.0-cap-CR-10626) (2022-12-30)

### Build

* use fixed docker repository because we can't reach accross jobs ([#2474](https://github.com/argoproj/argo-rollouts/issues/2474))
* copy proto files from GOPATH so we can clone outside of GOPATH ([#2360](https://github.com/argoproj/argo-rollouts/issues/2360))
* add sha256 checksums for all released bins ([#2332](https://github.com/argoproj/argo-rollouts/issues/2332))

### Chore

* remove deprecated -i for go build ([#2047](https://github.com/argoproj/argo-rollouts/issues/2047))
* Pin golang to 1.17 to avoid CVEs in docker image ([#1920](https://github.com/argoproj/argo-rollouts/issues/1920))
* Improve image build speed [#1919](https://github.com/argoproj/argo-rollouts/issues/1919) ([#1948](https://github.com/argoproj/argo-rollouts/issues/1948))
* Improve image build speed ([#1919](https://github.com/argoproj/argo-rollouts/issues/1919))
* update stable tag conditionally ([#2480](https://github.com/argoproj/argo-rollouts/issues/2480))
* fix checksum generation ([#2481](https://github.com/argoproj/argo-rollouts/issues/2481))
* add optum to users list ([#2466](https://github.com/argoproj/argo-rollouts/issues/2466))
* use docker login to sign images ([#2479](https://github.com/argoproj/argo-rollouts/issues/2479))
* use correct image for plugin container ([#2478](https://github.com/argoproj/argo-rollouts/issues/2478))
* release workflow docker build context should use local path and not git context ([#1388](https://github.com/argoproj/argo-rollouts/issues/1388))
* update version file to 1.2.0 ([#2013](https://github.com/argoproj/argo-rollouts/issues/2013))
* improve openapi schema ([#2081](https://github.com/argoproj/argo-rollouts/issues/2081))
* Add e2e and unit test comment reports ([#2123](https://github.com/argoproj/argo-rollouts/issues/2123))
* upgrade deps ([#2136](https://github.com/argoproj/argo-rollouts/issues/2136))
* Add Yotpo to USERS.md
* use controler-gen for cluster analysis template scope ([#2148](https://github.com/argoproj/argo-rollouts/issues/2148))
* Upgrade golang ([#2160](https://github.com/argoproj/argo-rollouts/issues/2160))
* Add example for istio-subset-split ([#2318](https://github.com/argoproj/argo-rollouts/issues/2318))
* upgrade golang to 1.19 ([#2219](https://github.com/argoproj/argo-rollouts/issues/2219))
* github release action was using incorect docker cache ([#1387](https://github.com/argoproj/argo-rollouts/issues/1387))
* sign container images and checksum assets ([#2334](https://github.com/argoproj/argo-rollouts/issues/2334))
* rename the examples/trafffic-management directory to istio ([#2315](https://github.com/argoproj/argo-rollouts/issues/2315))
* add deprecation notice for rollout_phase in docs ([#2377](https://github.com/argoproj/argo-rollouts/issues/2377)) ([#2378](https://github.com/argoproj/argo-rollouts/issues/2378))
* **cli:** add darwin arm64 to build and release ([#2264](https://github.com/argoproj/argo-rollouts/issues/2264))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.17.0 to 1.17.1 ([#2369](https://github.com/argoproj/argo-rollouts/issues/2369))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2417](https://github.com/argoproj/argo-rollouts/issues/2417))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2414](https://github.com/argoproj/argo-rollouts/issues/2414))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2413](https://github.com/argoproj/argo-rollouts/issues/2413))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2412](https://github.com/argoproj/argo-rollouts/issues/2412))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2409](https://github.com/argoproj/argo-rollouts/issues/2409))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2406](https://github.com/argoproj/argo-rollouts/issues/2406))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2404](https://github.com/argoproj/argo-rollouts/issues/2404))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2418](https://github.com/argoproj/argo-rollouts/issues/2418))
* **deps:** bump codecov/codecov-action from 2.1.0 to 3.1.1 ([#2251](https://github.com/argoproj/argo-rollouts/issues/2251))
* **deps:** bump google.golang.org/grpc from 1.50.1 to 1.51.0 ([#2421](https://github.com/argoproj/argo-rollouts/issues/2421))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2428](https://github.com/argoproj/argo-rollouts/issues/2428))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2429](https://github.com/argoproj/argo-rollouts/issues/2429))
* **deps:** bump dependabot/fetch-metadata from 1.3.4 to 1.3.5 ([#2390](https://github.com/argoproj/argo-rollouts/issues/2390))
* **deps:** bump imjasonh/setup-crane from 0.1 to 0.2 ([#2387](https://github.com/argoproj/argo-rollouts/issues/2387))
* **deps:** upgrade ui deps to fix high security cve's ([#2345](https://github.com/argoproj/argo-rollouts/issues/2345))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2430](https://github.com/argoproj/argo-rollouts/issues/2430))
* **deps:** bump actions/upload-artifact from 2 to 3 ([#1973](https://github.com/argoproj/argo-rollouts/issues/1973))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 ([#2381](https://github.com/argoproj/argo-rollouts/issues/2381))
* **deps:** bump github.com/spf13/cobra from 1.6.0 to 1.6.1 ([#2370](https://github.com/argoproj/argo-rollouts/issues/2370))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2366](https://github.com/argoproj/argo-rollouts/issues/2366))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2367](https://github.com/argoproj/argo-rollouts/issues/2367))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2439](https://github.com/argoproj/argo-rollouts/issues/2439))
* **deps:** bump github.com/stretchr/testify from 1.8.0 to 1.8.1 ([#2368](https://github.com/argoproj/argo-rollouts/issues/2368))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2365](https://github.com/argoproj/argo-rollouts/issues/2365))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.16.16 to 1.17.0 ([#2364](https://github.com/argoproj/argo-rollouts/issues/2364))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2361](https://github.com/argoproj/argo-rollouts/issues/2361))
* **deps:** bump github.com/prometheus/client_model from 0.2.0 to 0.3.0 ([#2349](https://github.com/argoproj/argo-rollouts/issues/2349))
* **deps:** bump github.com/valyala/fasttemplate from 1.2.1 to 1.2.2 ([#2348](https://github.com/argoproj/argo-rollouts/issues/2348))
* **deps:** bump github.com/newrelic/newrelic-client-go ([#2344](https://github.com/argoproj/argo-rollouts/issues/2344))
* **deps:** bump google.golang.org/grpc from 1.50.0 to 1.50.1 ([#2340](https://github.com/argoproj/argo-rollouts/issues/2340))
* **deps:** bump github.com/prometheus/common from 0.36.0 to 0.37.0 ([#2143](https://github.com/argoproj/argo-rollouts/issues/2143))
* **deps:** bump github.com/sirupsen/logrus from 1.8.1 to 1.9.0 ([#2152](https://github.com/argoproj/argo-rollouts/issues/2152))
* **deps:** bump github.com/spf13/cobra from 1.5.0 to 1.6.0 ([#2313](https://github.com/argoproj/argo-rollouts/issues/2313))
* **deps:** bump github.com/newrelic/newrelic-client-go ([#2267](https://github.com/argoproj/argo-rollouts/issues/2267))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2307](https://github.com/argoproj/argo-rollouts/issues/2307))
* **deps:** bump docker/build-push-action from 2 to 3 ([#2306](https://github.com/argoproj/argo-rollouts/issues/2306))
* **deps:** bump docker/setup-buildx-action from 1 to 2 ([#2305](https://github.com/argoproj/argo-rollouts/issues/2305))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 ([#2304](https://github.com/argoproj/argo-rollouts/issues/2304))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2295](https://github.com/argoproj/argo-rollouts/issues/2295))
* **deps:** bump google.golang.org/protobuf from 1.28.0 to 1.28.1 ([#2296](https://github.com/argoproj/argo-rollouts/issues/2296))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2255](https://github.com/argoproj/argo-rollouts/issues/2255))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2294](https://github.com/argoproj/argo-rollouts/issues/2294))
* **deps:** bump google.golang.org/grpc from 1.47.0 to 1.50.0 ([#2293](https://github.com/argoproj/argo-rollouts/issues/2293))
* **deps:** bump docker/metadata-action from 3 to 4 ([#2292](https://github.com/argoproj/argo-rollouts/issues/2292))
* **deps:** bump github/codeql-action from 1 to 2 ([#2289](https://github.com/argoproj/argo-rollouts/issues/2289))
* **deps:** bump docker/login-action from 1 to 2 ([#2288](https://github.com/argoproj/argo-rollouts/issues/2288))
* **deps:** bump actions/setup-go from 2 to 3 ([#2287](https://github.com/argoproj/argo-rollouts/issues/2287))
* **deps:** bump dependabot/fetch-metadata from 1.3.3 to 1.3.4 ([#2286](https://github.com/argoproj/argo-rollouts/issues/2286))
* **deps:** bump EnricoMi/publish-unit-test-result-action from 1 to 2 ([#2285](https://github.com/argoproj/argo-rollouts/issues/2285))
* **deps:** bump actions/setup-python from 2 to 4.1.0 ([#2134](https://github.com/argoproj/argo-rollouts/issues/2134))
* **deps:** bump actions/cache from 2 to 3.0.1 ([#1940](https://github.com/argoproj/argo-rollouts/issues/1940))
* **deps:** bump docker/setup-qemu-action from 1 to 2 ([#2284](https://github.com/argoproj/argo-rollouts/issues/2284))
* **deps:** bump actions/checkout from 2 to 3.1.0 ([#2283](https://github.com/argoproj/argo-rollouts/issues/2283))
* **deps:** bump github.com/influxdata/influxdb-client-go/v2 ([#2447](https://github.com/argoproj/argo-rollouts/issues/2447))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2452](https://github.com/argoproj/argo-rollouts/issues/2452))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2454](https://github.com/argoproj/argo-rollouts/issues/2454))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2455](https://github.com/argoproj/argo-rollouts/issues/2455))
* **deps:** bump github.com/prometheus/common from 0.37.0 to 0.38.0 ([#2468](https://github.com/argoproj/argo-rollouts/issues/2468))
* **deps:** bump github.com/prometheus/client_golang ([#2469](https://github.com/argoproj/argo-rollouts/issues/2469))
* **deps:** bump notification engine ([#2470](https://github.com/argoproj/argo-rollouts/issues/2470))
* **deps:** bump github.com/prometheus/common from 0.38.0 to 0.39.0 ([#2476](https://github.com/argoproj/argo-rollouts/issues/2476))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2477](https://github.com/argoproj/argo-rollouts/issues/2477))
* **deps:** bump github.com/aws/aws-sdk-go-v2 from 1.17.2 to 1.17.3 ([#2484](https://github.com/argoproj/argo-rollouts/issues/2484))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#2486](https://github.com/argoproj/argo-rollouts/issues/2486))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#2485](https://github.com/argoproj/argo-rollouts/issues/2485))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/cloudwatch ([#2487](https://github.com/argoproj/argo-rollouts/issues/2487))
* **doc:** Clarify doc for Rollout.spec.progressDeadlineAbort Signed-off-by: Vladimir Ivanov <vladee.ivanov[@gmail](https://github.com/gmail).com>
* **doc:** Sync doc and code comments for the progressDeadlineAbort param

### Ci

* create stable tag for docs ([#2402](https://github.com/argoproj/argo-rollouts/issues/2402))
* fix some github actions warnings ([#2375](https://github.com/argoproj/argo-rollouts/issues/2375))
* add link to conventional pr check in pr template ([#2346](https://github.com/argoproj/argo-rollouts/issues/2346))
* auto generate changelog ([#2321](https://github.com/argoproj/argo-rollouts/issues/2321))
* adjust settings for stale pr and issues ([#2341](https://github.com/argoproj/argo-rollouts/issues/2341))
* fix pr lint check ([#2336](https://github.com/argoproj/argo-rollouts/issues/2336))
* add auto close to issues and prs ([#2319](https://github.com/argoproj/argo-rollouts/issues/2319))
* Add github action for PR Conventional Commits ([#2320](https://github.com/argoproj/argo-rollouts/issues/2320))
* Auto-cancel redundant builds. ([#2002](https://github.com/argoproj/argo-rollouts/issues/2002))

### Cleanup

* rename temlateref to templateref ([#2154](https://github.com/argoproj/argo-rollouts/issues/2154))

### Docs

* fix rendering by upgrading deps ([#2495](https://github.com/argoproj/argo-rollouts/issues/2495))
* Add traffic router support to readme ([#2444](https://github.com/argoproj/argo-rollouts/issues/2444))
* fix typo in helm Argo rollouts ([#2442](https://github.com/argoproj/argo-rollouts/issues/2442))
* correct syntax of canary setMirrorRoute's value ([#2431](https://github.com/argoproj/argo-rollouts/issues/2431))
* Explain upgrade process ([#2424](https://github.com/argoproj/argo-rollouts/issues/2424))
* add progressive delivery with gitops example for openshift ([#2400](https://github.com/argoproj/argo-rollouts/issues/2400))
* fix !important block typo ([#2372](https://github.com/argoproj/argo-rollouts/issues/2372))
* mention supported versions ([#2163](https://github.com/argoproj/argo-rollouts/issues/2163))
* Added blog post for minimize impact in Kubernetes using Progressive Delivery and customer side impact ([#2355](https://github.com/argoproj/argo-rollouts/issues/2355))
* steps to rollback to deployment kinds ([#2014](https://github.com/argoproj/argo-rollouts/issues/2014))
* add artifact badge ([#2331](https://github.com/argoproj/argo-rollouts/issues/2331))
* Use new Google Analytics 4 site tag ([#2299](https://github.com/argoproj/argo-rollouts/issues/2299))
* Fixed read the docs rendering ([#2277](https://github.com/argoproj/argo-rollouts/issues/2277))
* common questions for Rollbacks ([#2027](https://github.com/argoproj/argo-rollouts/issues/2027))
* add OpsVerse as an official user (USERS.md) ([#2209](https://github.com/argoproj/argo-rollouts/issues/2209))
* Fix the controller annotation to enable data scrapping ([#2238](https://github.com/argoproj/argo-rollouts/issues/2238))
* Update release docs for versioned formula ([#2245](https://github.com/argoproj/argo-rollouts/issues/2245))
* Update docs for new openapi kustomize support ([#2216](https://github.com/argoproj/argo-rollouts/issues/2216))
* add Opensurvey to USERS.md ([#2195](https://github.com/argoproj/argo-rollouts/issues/2195))
* update release doc with brew formula details ([#2165](https://github.com/argoproj/argo-rollouts/issues/2165))
* add selector to migrating page ([#2039](https://github.com/argoproj/argo-rollouts/issues/2039))
* expose/fix Traefik docs ([#2017](https://github.com/argoproj/argo-rollouts/issues/2017))
* add WorkloadRef to Rollout spec ([#2019](https://github.com/argoproj/argo-rollouts/issues/2019))
* **trafficrouting:** fix docs warning to github style markdown ([#2342](https://github.com/argoproj/argo-rollouts/issues/2342))

### Feat

* Ability for lint command to inspect referenced resources ([#2030](https://github.com/argoproj/argo-rollouts/issues/2030))
* Apache APISIX support. Fixes [#2395](https://github.com/argoproj/argo-rollouts/issues/2395) ([#2437](https://github.com/argoproj/argo-rollouts/issues/2437))
* rollback windows. Fixes [#574](https://github.com/argoproj/argo-rollouts/issues/574) ([#2394](https://github.com/argoproj/argo-rollouts/issues/2394))
* Report notification metrics for rollouts ([#1856](https://github.com/argoproj/argo-rollouts/issues/1856))
* Allow prometheus server address to be centrally configured ([#1956](https://github.com/argoproj/argo-rollouts/issues/1956))
* Traefik support. Fixes [#516](https://github.com/argoproj/argo-rollouts/issues/516) ([#1907](https://github.com/argoproj/argo-rollouts/issues/1907))
* add support for getting the replicaset name via templating ([#2396](https://github.com/argoproj/argo-rollouts/issues/2396))
* Allow Traffic shaping through header based routing for ALB ([#2214](https://github.com/argoproj/argo-rollouts/issues/2214))
* Add support for spec.ingressClassName ([#2178](https://github.com/argoproj/argo-rollouts/issues/2178))
* Support TCP routes traffic splitting for Istio VirtualService ([#1659](https://github.com/argoproj/argo-rollouts/issues/1659))
* emit rollout delete event ([#1893](https://github.com/argoproj/argo-rollouts/issues/1893))
* support /rollouts/:namespace?q=... and /rollout/:namespace/:name ([#1902](https://github.com/argoproj/argo-rollouts/issues/1902))
* Allow Traffic shaping through header based routing. Fixes [#474](https://github.com/argoproj/argo-rollouts/issues/474) ([#1990](https://github.com/argoproj/argo-rollouts/issues/1990))
* Adds support for Istio traffic mirroring ([#2074](https://github.com/argoproj/argo-rollouts/issues/2074))
* add support for influxdb as a metrics provider ([#1839](https://github.com/argoproj/argo-rollouts/issues/1839))
* ArgoRollouts dashboard now supporting rootpath ([#2075](https://github.com/argoproj/argo-rollouts/issues/2075))
* Implement Issue [#1779](https://github.com/argoproj/argo-rollouts/issues/1779): add rollout.Spec.Strategy.Canary.MinPodsPerReplicaSet ([#2448](https://github.com/argoproj/argo-rollouts/issues/2448))
* support Ingress from Networking API version ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529))
* **cli:** add port flag for dashboard command ([#2383](https://github.com/argoproj/argo-rollouts/issues/2383))
* **cli:** dynamic shell completion for main resources names (rollouts, experiments, analysisrun) ([#2379](https://github.com/argoproj/argo-rollouts/issues/2379))
* **controller:** don't hardcode experiment ports; always create service ([#2397](https://github.com/argoproj/argo-rollouts/issues/2397))
* **grafana:** Allow selecting datasource for grafana dashboard ([#1988](https://github.com/argoproj/argo-rollouts/issues/1988))

### Feature

* Dashboard now displaying analysis details ([#1910](https://github.com/argoproj/argo-rollouts/issues/1910))
* Dashboard displaying the setCanaryScale values ([#1923](https://github.com/argoproj/argo-rollouts/issues/1923))

### Fix

* e2e istio crd; deprecated apiextensions/v1beta1 ([#1740](https://github.com/argoproj/argo-rollouts/issues/1740))
* dev build can set DEV_IMAGE=true ([#2440](https://github.com/argoproj/argo-rollouts/issues/2440))
* add patch verb to deployment resource ([#2407](https://github.com/argoproj/argo-rollouts/issues/2407))
* unsolicited rollout after upgrade from v0.10->v1.0 when pod was using service account ([#1367](https://github.com/argoproj/argo-rollouts/issues/1367))
* Abort rollout doesn't remove all canary pods for setCanaryScale ([#1352](https://github.com/argoproj/argo-rollouts/issues/1352))
* UI not redirecting on / ([#2252](https://github.com/argoproj/argo-rollouts/issues/2252))
* set gopath in makefile ([#2398](https://github.com/argoproj/argo-rollouts/issues/2398))
* nil pointer dereference when reconciling paused blue-green rollout ([#1378](https://github.com/argoproj/argo-rollouts/issues/1378))
* Promote full did not work against BlueGreen with previewReplicaCount ([#1384](https://github.com/argoproj/argo-rollouts/issues/1384))
* canary scaledown event could violate maxUnavailable ([#1429](https://github.com/argoproj/argo-rollouts/issues/1429))
* analysis runs to wait for all metrics to complete ([#1407](https://github.com/argoproj/argo-rollouts/issues/1407))
* nil pointer while linting with basic canary and ingresses ([#2256](https://github.com/argoproj/argo-rollouts/issues/2256))
* retarget blue-green previewService before scaling up preview ReplicaSet ([#1368](https://github.com/argoproj/argo-rollouts/issues/1368))
* Analysis argument validation ([#1412](https://github.com/argoproj/argo-rollouts/issues/1412))
* enable notifications without when condition ([#2231](https://github.com/argoproj/argo-rollouts/issues/2231))
* change completed condition so it only triggers on pod hash changes also adds an event for when it  does changes. ([#2203](https://github.com/argoproj/argo-rollouts/issues/2203))
* rootPath support so that it uses the embedded files system ([#2198](https://github.com/argoproj/argo-rollouts/issues/2198))
* Nginx ingressClassName passed to canary ingress ([#1448](https://github.com/argoproj/argo-rollouts/issues/1448))
* Failed to process: Object 'Kind' is missing in Errors with rollouts notification ([#2150](https://github.com/argoproj/argo-rollouts/issues/2150))
* replica count for new deployment ([#1449](https://github.com/argoproj/argo-rollouts/issues/1449))
* remove metrics when objects are removed from cluster to prevent build up ([#2115](https://github.com/argoproj/argo-rollouts/issues/2115))
* Update ro.Status.ALB when first creating rollout object ([#1986](https://github.com/argoproj/argo-rollouts/issues/1986))
* argo-rollouts occasionally crashes in argoproj/pkg ([#2111](https://github.com/argoproj/argo-rollouts/issues/2111))
* High reconciliation activity and CPU load for invalid rollout ([#2091](https://github.com/argoproj/argo-rollouts/issues/2091))
* notifications when condition ([#2066](https://github.com/argoproj/argo-rollouts/issues/2066))
* UI codegen ([#2072](https://github.com/argoproj/argo-rollouts/issues/2072))
* missing lb event ([#2021](https://github.com/argoproj/argo-rollouts/issues/2021))
* Change behavior of rollout to not check for availability during rollout and fix flakey e2e tests ([#1957](https://github.com/argoproj/argo-rollouts/issues/1957))
* Add pagination to FindLoadBalancerByDNSName ([#1971](https://github.com/argoproj/argo-rollouts/issues/1971))
* Add watch verb to clusterRole for pods
* build/lint is broken due to dependencies changes ([#1958](https://github.com/argoproj/argo-rollouts/issues/1958))
* Use actual weight from status field on rollout object ([#1937](https://github.com/argoproj/argo-rollouts/issues/1937))
* Handle minor version with '+' when determining ingress mode ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529)) ([#1612](https://github.com/argoproj/argo-rollouts/issues/1612))
* nginx traffic router patching wrong ingress resource ([#1655](https://github.com/argoproj/argo-rollouts/issues/1655))
* default replica before resolving workloadRef ([#1304](https://github.com/argoproj/argo-rollouts/issues/1304))
* **analysis:** Fix Analysis Terminal Decision For Dry-Run Metrics ([#2131](https://github.com/argoproj/argo-rollouts/issues/2131))
* **analysis:** Avoid Infinite Error Message Append For Failed Dry-Run Metrics ([#2182](https://github.com/argoproj/argo-rollouts/issues/2182))
* **analysis:** Make AR End When Only Dry-Run Metrics Are Defined ([#2230](https://github.com/argoproj/argo-rollouts/issues/2230))
* **analysis:** Fix Analysis Terminal Decision For Dry-Run Metrics ([#2399](https://github.com/argoproj/argo-rollouts/issues/2399))
* **cli:** nil pointer while linting  ([#2324](https://github.com/argoproj/argo-rollouts/issues/2324))
* **controller:** leader election preventing two controllers running and gracefully shutting down ([#2291](https://github.com/argoproj/argo-rollouts/issues/2291))
* **controller:**  Fix k8s clientset controller metrics. Fixes [#2139](https://github.com/argoproj/argo-rollouts/issues/2139) ([#2261](https://github.com/argoproj/argo-rollouts/issues/2261))
* **dashboard:** correct mime type is returned. Fixes: [#2290](https://github.com/argoproj/argo-rollouts/issues/2290) ([#2303](https://github.com/argoproj/argo-rollouts/issues/2303))
* **e2e:** DeFlake E2E Tests [#1647](https://github.com/argoproj/argo-rollouts/issues/1647) ([#1648](https://github.com/argoproj/argo-rollouts/issues/1648))
* **example:** correct docs when metrics got result empty ([#2309](https://github.com/argoproj/argo-rollouts/issues/2309))
* **metricprovider:** Support jsonBody for web metric provider Fixes [#2275](https://github.com/argoproj/argo-rollouts/issues/2275) ([#2312](https://github.com/argoproj/argo-rollouts/issues/2312))
* **trafficrouting:** Do not block the switch of service selectors for single pod failures ([#2441](https://github.com/argoproj/argo-rollouts/issues/2441))

### Fixes

* **controller:** istio dropping fields not defined in type ([#2268](https://github.com/argoproj/argo-rollouts/issues/2268))

### Test

* **controller:** add extra checks to TestWriteBackToInformer ([#2326](https://github.com/argoproj/argo-rollouts/issues/2326))


<a name="v1.2.0"></a>
## [v1.2.0](https://github.com/argoproj/argo-rollouts/compare/v1.2.0-rc2...v1.2.0) (2022-03-21)

### Chore

* fix golangci-lint to 1.44 to fix build error ([#1917](https://github.com/argoproj/argo-rollouts/issues/1917))
* move dependencies to dev dependencies

### Docs

* vpa for rollouts ([#1909](https://github.com/argoproj/argo-rollouts/issues/1909))
* Add SAP Concur ([#1878](https://github.com/argoproj/argo-rollouts/issues/1878))

### Feat

* Dashboard now displaying name, specRef and weight in the experimental step. ([#1863](https://github.com/argoproj/argo-rollouts/issues/1863))


<a name="v1.2.0-rc2"></a>
## [v1.2.0-rc2](https://github.com/argoproj/argo-rollouts/compare/v1.2.0-rc1...v1.2.0-rc2) (2022-02-17)

### Docs

* add community section to README.md
* Fix codegen of mkdocs.yaml
* Fix lint
* Add new line
* Add a toggle for dark mode

### Feat

* Added delay button in the scaled down revision ([#1355](https://github.com/argoproj/argo-rollouts/issues/1355)) ([#1804](https://github.com/argoproj/argo-rollouts/issues/1804))

### Fix

* add workaround to fix 'stream terminated by RST_STREAM with error code: PROTOCOL_ERROR' ([#1862](https://github.com/argoproj/argo-rollouts/issues/1862))


<a name="v1.2.0-rc1"></a>
## [v1.2.0-rc1](https://github.com/argoproj/argo-rollouts/compare/v1.2.0-cap-CR-10626...v1.2.0-rc1) (2022-02-05)


<a name="v1.2.0-cap-CR-10626"></a>
## [v1.2.0-cap-CR-10626](https://github.com/argoproj/argo-rollouts/compare/v1.1.0...v1.2.0-cap-CR-10626) (2022-05-12)

### Chore

* upgrade k8s libraries to v1.22 ([#1773](https://github.com/argoproj/argo-rollouts/issues/1773))
* move dependencies to dev dependencies
* fix spdx image generation ([#1849](https://github.com/argoproj/argo-rollouts/issues/1849))
* fix spdx ci ([#1848](https://github.com/argoproj/argo-rollouts/issues/1848))
* Generate spdx file for the docker image ([#1844](https://github.com/argoproj/argo-rollouts/issues/1844))
* release workflow docker build context should use local path and not git context ([#1388](https://github.com/argoproj/argo-rollouts/issues/1388))
* update k8s pkg to resolve vulnerabilities  ([#1545](https://github.com/argoproj/argo-rollouts/issues/1545))
* CVE-2020-26160 ([#1829](https://github.com/argoproj/argo-rollouts/issues/1829))
* generate and upload sbom during release ([#1834](https://github.com/argoproj/argo-rollouts/issues/1834))
* pin sys module in go.mod to resolve a fatal runtime execution in go 1.17 ([#1692](https://github.com/argoproj/argo-rollouts/issues/1692))
* fix flaky TestAbortRolloutAfterFailedExperiment test ([#1710](https://github.com/argoproj/argo-rollouts/issues/1710))
* update docs for minikube 1.19 ([#1746](https://github.com/argoproj/argo-rollouts/issues/1746))
* Configure dependabot to ignore k8s dependencies ([#1802](https://github.com/argoproj/argo-rollouts/issues/1802))
* fix golangci-lint to 1.44 to fix build error ([#1917](https://github.com/argoproj/argo-rollouts/issues/1917))
* Fix istio vs reconcile errors ([#1460](https://github.com/argoproj/argo-rollouts/issues/1460))
* make ci/local codegen consistent ([#1772](https://github.com/argoproj/argo-rollouts/issues/1772))
* github release action was using incorect docker cache ([#1387](https://github.com/argoproj/argo-rollouts/issues/1387))
* **deps:** bump github.com/antonmedv/expr from 1.8.9 to 1.9.0 ([#1712](https://github.com/argoproj/argo-rollouts/issues/1712))
* **deps:** update github.com/miekg/dns for CVE-2019-19794 ([#1810](https://github.com/argoproj/argo-rollouts/issues/1810))
* **deps:** bump codecov/codecov-action from 2.0.3 to 2.1.0 ([#1508](https://github.com/argoproj/argo-rollouts/issues/1508))
* **deps:** bump github.com/evanphx/json-patch/v5 from 5.2.0 to 5.6.0 ([#1603](https://github.com/argoproj/argo-rollouts/issues/1603))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#1797](https://github.com/argoproj/argo-rollouts/issues/1797))
* **deps:** bump github.com/prometheus/common from 0.21.0 to 0.32.1 ([#1604](https://github.com/argoproj/argo-rollouts/issues/1604))
* **deps:** bump github.com/aws/aws-sdk-go-v2 dependencies ([#1835](https://github.com/argoproj/argo-rollouts/issues/1835))
* **deps:** bump github.com/newrelic/newrelic-client-go ([#1836](https://github.com/argoproj/argo-rollouts/issues/1836))
* **deps:** bump github.com/aws/aws-sdk-go-v2/config ([#1791](https://github.com/argoproj/argo-rollouts/issues/1791))
* **docs:** Updated FAQs ([#1695](https://github.com/argoproj/argo-rollouts/issues/1695))

### Docs

* vpa for rollouts ([#1909](https://github.com/argoproj/argo-rollouts/issues/1909))
* Add SAP Concur ([#1878](https://github.com/argoproj/argo-rollouts/issues/1878))
* add community section to README.md
* Fix codegen of mkdocs.yaml
* Fix lint
* Add new line
* Add a toggle for dark mode
* Updated the case of properties according to actual property name
* Update security.md ([#1840](https://github.com/argoproj/argo-rollouts/issues/1840))
* fixed rollout controller link
* mention internal architecture
* Added ArgoCon 21 presentation ([#1811](https://github.com/argoproj/argo-rollouts/issues/1811))
* update membership information ([#1814](https://github.com/argoproj/argo-rollouts/issues/1814))
* fix doc for valueFrom fields for analysis args ([#1763](https://github.com/argoproj/argo-rollouts/issues/1763))
* Add Gllue to list of users ([#1745](https://github.com/argoproj/argo-rollouts/issues/1745))
* Add Ibotta to the list of users ([#1744](https://github.com/argoproj/argo-rollouts/issues/1744))
* mention default notification templates ([#1725](https://github.com/argoproj/argo-rollouts/issues/1725))
* Add notiication templates for rollouts and analysis ([#1753](https://github.com/argoproj/argo-rollouts/issues/1753))
* clarify the setCanaryScale of dynamic canary scale ([#1703](https://github.com/argoproj/argo-rollouts/issues/1703))
* Added readthedocs configuration ([#1528](https://github.com/argoproj/argo-rollouts/issues/1528))
* Use readthedocs versionining. Closes [#1518](https://github.com/argoproj/argo-rollouts/issues/1518) ([#1671](https://github.com/argoproj/argo-rollouts/issues/1671))
* fix some vague description about analysis arguments ([#1672](https://github.com/argoproj/argo-rollouts/issues/1672))
* Clarify application dependencies ([#1706](https://github.com/argoproj/argo-rollouts/issues/1706))
* Add Akuity to the list of users ([#1598](https://github.com/argoproj/argo-rollouts/issues/1598))
* Add link to awesome-argo for more resources ([#1622](https://github.com/argoproj/argo-rollouts/issues/1622))
* example data is used in both examples ([#1570](https://github.com/argoproj/argo-rollouts/issues/1570))
* **analysis:** add missing explanation about failureLimit ([#1674](https://github.com/argoproj/argo-rollouts/issues/1674))
* **controller:** remove duplicate sentence. ([#1756](https://github.com/argoproj/argo-rollouts/issues/1756))

### Feat

* increase default QPS/Burst to 40/80. Allow values to be tunable ([#1679](https://github.com/argoproj/argo-rollouts/issues/1679))
* support Ingress from Networking API version ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529))
* Added delay button in the scaled down revision ([#1355](https://github.com/argoproj/argo-rollouts/issues/1355)) ([#1804](https://github.com/argoproj/argo-rollouts/issues/1804))
* Add logformat flag to rollouts-controller ([#1818](https://github.com/argoproj/argo-rollouts/issues/1818))
* ping-pong service management ([#1697](https://github.com/argoproj/argo-rollouts/issues/1697))
* TrafficRouting support with AWS App Mesh ([#1401](https://github.com/argoproj/argo-rollouts/issues/1401)) ([#1606](https://github.com/argoproj/argo-rollouts/issues/1606))
* Istio Host-level TrafficRouting with experiment step ([#1569](https://github.com/argoproj/argo-rollouts/issues/1569))
* support Ingress from Networking API version ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529))
* add healthz probe port and update the install.yaml ([#1578](https://github.com/argoproj/argo-rollouts/issues/1578))
* Istio Subset-Level TrafficRouting with experiment step ([#1602](https://github.com/argoproj/argo-rollouts/issues/1602))
* added Argo version info in /metrics endpoint ([#1662](https://github.com/argoproj/argo-rollouts/issues/1662))
* Dashboard now displaying name, specRef and weight in the experimental step. ([#1863](https://github.com/argoproj/argo-rollouts/issues/1863))
* HA Leader election support on rollouts-controller ([#1519](https://github.com/argoproj/argo-rollouts/issues/1519))
* **alb:** Surface ALB information into rollout status ([#1241](https://github.com/argoproj/argo-rollouts/issues/1241)) ([#1625](https://github.com/argoproj/argo-rollouts/issues/1625))
* **analysis:** Add Dry-Run Mode ([#1627](https://github.com/argoproj/argo-rollouts/issues/1627))
* **analysis:** Add Measurements Retention Limit Option for Metrics ([#1729](https://github.com/argoproj/argo-rollouts/issues/1729))
* **analysis:** Allow analysis arguments to get valueFrom Rollout status ([#1242](https://github.com/argoproj/argo-rollouts/issues/1242)) ([#1629](https://github.com/argoproj/argo-rollouts/issues/1629))
* **analysis:** Added additional metadata to the status of AnalysisRun
* **controller:** multiple TrafficRoutingReconciler ([#1472](https://github.com/argoproj/argo-rollouts/issues/1472))
* **experiment:** Added DryRun analysis mode functionality for experiments ([#1691](https://github.com/argoproj/argo-rollouts/issues/1691))
* **experiment:** Add Measurements Retention Limit Option for Metrics
* **manifests:** Add Age column to all CRD(s). Fixes [#1511](https://github.com/argoproj/argo-rollouts/issues/1511) ([#1527](https://github.com/argoproj/argo-rollouts/issues/1527))
* **rollout:** AnalysisRuns created by Rollouts can limit retention of metrics ([#1780](https://github.com/argoproj/argo-rollouts/issues/1780))
* **webmetric:** Support POST/PUT content with web metrics. Fixes [#371](https://github.com/argoproj/argo-rollouts/issues/371) ([#1573](https://github.com/argoproj/argo-rollouts/issues/1573))

### Fix

* Handle minor version with '+' when determining ingress mode ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529)) ([#1612](https://github.com/argoproj/argo-rollouts/issues/1612))
* nginx traffic router patching wrong ingress resource ([#1655](https://github.com/argoproj/argo-rollouts/issues/1655))
* unsolicited rollout after upgrade from v0.10->v1.0 when pod was using service account ([#1367](https://github.com/argoproj/argo-rollouts/issues/1367))
* e2e istio crd; deprecated apiextensions/v1beta1 ([#1740](https://github.com/argoproj/argo-rollouts/issues/1740))
* add workaround to fix 'stream terminated by RST_STREAM with error code: PROTOCOL_ERROR' ([#1862](https://github.com/argoproj/argo-rollouts/issues/1862))
* Abort rollout doesn't remove all canary pods for setCanaryScale ([#1352](https://github.com/argoproj/argo-rollouts/issues/1352))
* using our own pod template hashing ([#1809](https://github.com/argoproj/argo-rollouts/issues/1809))
* flaky unit test ([#1831](https://github.com/argoproj/argo-rollouts/issues/1831))
* nil pointer dereference when reconciling paused blue-green rollout ([#1378](https://github.com/argoproj/argo-rollouts/issues/1378))
* client can detect if rollout is in the process of unpausing ([#1798](https://github.com/argoproj/argo-rollouts/issues/1798))
* remove non-existent target in makefile ([#1813](https://github.com/argoproj/argo-rollouts/issues/1813))
* canary replicas/weight could flap during abort with dynamic scaling ([#1794](https://github.com/argoproj/argo-rollouts/issues/1794))
* Promote-full with dynamicStableScaling increases weight according to available canary pods. Fixes: [#1681](https://github.com/argoproj/argo-rollouts/issues/1681) ([#1683](https://github.com/argoproj/argo-rollouts/issues/1683))
* plugin panic while watching progress ([#1796](https://github.com/argoproj/argo-rollouts/issues/1796))
* Promote full did not work against BlueGreen with previewReplicaCount ([#1384](https://github.com/argoproj/argo-rollouts/issues/1384))
* canary scaledown event could violate maxUnavailable ([#1429](https://github.com/argoproj/argo-rollouts/issues/1429))
* analysis runs to wait for all metrics to complete ([#1407](https://github.com/argoproj/argo-rollouts/issues/1407))
* plugin did not set deployment image when using workloadRef ([#1787](https://github.com/argoproj/argo-rollouts/issues/1787))
* notifications using workloadRef did not have access to pod template ([#1786](https://github.com/argoproj/argo-rollouts/issues/1786))
* delay service injection of selector labels until ReplicaSet available ([#1777](https://github.com/argoproj/argo-rollouts/issues/1777))
* controller could panic in scaling events with analysis ([#1699](https://github.com/argoproj/argo-rollouts/issues/1699))
* retry Experiment ReplicaSet scaling conflict errors ([#1778](https://github.com/argoproj/argo-rollouts/issues/1778))
* continue update process in middle of update if spec.replicas is 0 ([#1764](https://github.com/argoproj/argo-rollouts/issues/1764))
* status.alb should be optionally populated ([#1766](https://github.com/argoproj/argo-rollouts/issues/1766))
* missing array type in the CRD rollout's spec volumes ([#1737](https://github.com/argoproj/argo-rollouts/issues/1737))
* traffic routed canary would flap traffic to stable after last step ([#1757](https://github.com/argoproj/argo-rollouts/issues/1757))
* retarget blue-green previewService before scaling up preview ReplicaSet ([#1368](https://github.com/argoproj/argo-rollouts/issues/1368))
* e2e istio crd; deprecated apiextensions/v1beta1 ([#1740](https://github.com/argoproj/argo-rollouts/issues/1740))
* dashboard promote buttons disabled during deploy ([#1669](https://github.com/argoproj/argo-rollouts/issues/1669))
* leaderelection uses the lock in the same ns as the controller ([#1717](https://github.com/argoproj/argo-rollouts/issues/1717))
* release the dashboard-install.yaml ([#1601](https://github.com/argoproj/argo-rollouts/issues/1601))
* Analysis argument validation ([#1412](https://github.com/argoproj/argo-rollouts/issues/1412))
* missing rollout informer writeback ([#1698](https://github.com/argoproj/argo-rollouts/issues/1698))
* use patch to update workload-generation annotation ([#1678](https://github.com/argoproj/argo-rollouts/issues/1678))
* sending updates to dashboard when a pod terminates ([#1642](https://github.com/argoproj/argo-rollouts/issues/1642))
* reset the progress condition when a pod is restarted ([#1649](https://github.com/argoproj/argo-rollouts/issues/1649))
* Fixed NPE while getting the ReplicaSet labels ([#1664](https://github.com/argoproj/argo-rollouts/issues/1664))
* add service delete to argo-rollouts role ([#1632](https://github.com/argoproj/argo-rollouts/issues/1632))
* Modify Experiment collision naming from dot-notation to dash ([#1646](https://github.com/argoproj/argo-rollouts/issues/1646))
* Wait for all canary pods to come up in TrafficRouting canary before switching traffic ([#1663](https://github.com/argoproj/argo-rollouts/issues/1663))
* Route traffic to Experiment even if Canary RS not scaled ([#1638](https://github.com/argoproj/argo-rollouts/issues/1638))
* Nginx ingressClassName passed to canary ingress ([#1448](https://github.com/argoproj/argo-rollouts/issues/1448))
* nginx traffic router patching wrong ingress resource ([#1655](https://github.com/argoproj/argo-rollouts/issues/1655))
* replica count for new deployment ([#1449](https://github.com/argoproj/argo-rollouts/issues/1449))
* validate service selctor labels matching rollout template labels ([#1618](https://github.com/argoproj/argo-rollouts/issues/1618))
* inconsistent status command output ([#1433](https://github.com/argoproj/argo-rollouts/issues/1433))
* rollout experiment template changing reference rs template labels. Fixes [#1596](https://github.com/argoproj/argo-rollouts/issues/1596) ([#1597](https://github.com/argoproj/argo-rollouts/issues/1597))
* Handle minor version with '+' when determining ingress mode ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529)) ([#1612](https://github.com/argoproj/argo-rollouts/issues/1612))
* Enable default triggers for argo rollouts ([#1689](https://github.com/argoproj/argo-rollouts/issues/1689))
* viewcontroller gorouting leak in status and get subcommand ([#1584](https://github.com/argoproj/argo-rollouts/issues/1584))
* default replica before resolving workloadRef ([#1304](https://github.com/argoproj/argo-rollouts/issues/1304))
* **analysis:** surface analysis validation failure to rollout status ([#1833](https://github.com/argoproj/argo-rollouts/issues/1833))
* **canary:** scale up and down old replicas ([#1824](https://github.com/argoproj/argo-rollouts/issues/1824))
* **controller:** Sticky session correction for AWS ALB. Fixes [#1572](https://github.com/argoproj/argo-rollouts/issues/1572) ([#1577](https://github.com/argoproj/argo-rollouts/issues/1577))
* **docs:** Remove Non-Existent Metrics From Docs ([#1650](https://github.com/argoproj/argo-rollouts/issues/1650))
* **e2e:** DeFlake E2E Tests [#1647](https://github.com/argoproj/argo-rollouts/issues/1647) ([#1648](https://github.com/argoproj/argo-rollouts/issues/1648))
* **e2e:** DeFlake E2E Tests [#1647](https://github.com/argoproj/argo-rollouts/issues/1647) ([#1648](https://github.com/argoproj/argo-rollouts/issues/1648))
* **plugin:** Fixes arm64 compatibility to plugin docker image. Fixes [#1728](https://github.com/argoproj/argo-rollouts/issues/1728) ([#1732](https://github.com/argoproj/argo-rollouts/issues/1732))
* **ui:** Show container images in dashboard for rollouts with a WorkloadRef ([#1792](https://github.com/argoproj/argo-rollouts/issues/1792))
* **ui:** Requesting cluster scoped namespaces does not fail gracefully in UI ([#1795](https://github.com/argoproj/argo-rollouts/issues/1795))
* **ui:** Truncate long container names ([#1793](https://github.com/argoproj/argo-rollouts/issues/1793))

### Refactor

* stop using mpatch ([#1654](https://github.com/argoproj/argo-rollouts/issues/1654))


<a name="v1.1.0"></a>
## [v1.1.0](https://github.com/argoproj/argo-rollouts/compare/v1.1.0-rc2...v1.1.0) (2021-10-11)

### Feat

* add default() evaluate helper. allow empty datadog result. Fixes [#1548](https://github.com/argoproj/argo-rollouts/issues/1548) ([#1551](https://github.com/argoproj/argo-rollouts/issues/1551))

### Fix

* change virtualService to pointer ([#1558](https://github.com/argoproj/argo-rollouts/issues/1558))


<a name="v1.1.0-rc2"></a>
## [v1.1.0-rc2](https://github.com/argoproj/argo-rollouts/compare/v1.1.0-rc1...v1.1.0-rc2) (2021-09-29)

### Chore

* publish notifications-install.yaml and rollout_cr_schema.json as part of release ([#1532](https://github.com/argoproj/argo-rollouts/issues/1532))
* temporarily disable ghcr.io image pushing ([#1530](https://github.com/argoproj/argo-rollouts/issues/1530))

### Docs

* update CHANGELOG.md for v1.1 release ([#1533](https://github.com/argoproj/argo-rollouts/issues/1533))
* clarify the service in notificatioin.md ([#1546](https://github.com/argoproj/argo-rollouts/issues/1546))

### Fix

* Istio does not switch the traffic. Fix the VS new object creation when the tls routes nil ([#1553](https://github.com/argoproj/argo-rollouts/issues/1553))
* **metricproviders:** Check and handle invalid server URL. Fixed [#1444](https://github.com/argoproj/argo-rollouts/issues/1444) ([#1534](https://github.com/argoproj/argo-rollouts/issues/1534))

### Fix

* Failed analysis to degrade rollout when multiple metrics are analyzed ([#1535](https://github.com/argoproj/argo-rollouts/issues/1535))


<a name="v1.1.0-rc1"></a>
## [v1.1.0-rc1](https://github.com/argoproj/argo-rollouts/compare/v1.1.0-cap-CR-8836...v1.1.0-rc1) (2021-09-21)


<a name="v1.1.0-cap-CR-8836"></a>
## [v1.1.0-cap-CR-8836](https://github.com/argoproj/argo-rollouts/compare/v1.1.0-cap-CR-7557...v1.1.0-cap-CR-8836) (2022-03-09)

### Feat

* support Ingress from Networking API version ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529))

### Fix

* Handle minor version with '+' when determining ingress mode ([#1529](https://github.com/argoproj/argo-rollouts/issues/1529)) ([#1612](https://github.com/argoproj/argo-rollouts/issues/1612))
* nginx traffic router patching wrong ingress resource ([#1655](https://github.com/argoproj/argo-rollouts/issues/1655))
* e2e istio crd; deprecated apiextensions/v1beta1 ([#1740](https://github.com/argoproj/argo-rollouts/issues/1740))
* **e2e:** DeFlake E2E Tests [#1647](https://github.com/argoproj/argo-rollouts/issues/1647) ([#1648](https://github.com/argoproj/argo-rollouts/issues/1648))


<a name="v1.1.0-cap-CR-7557"></a>
## [v1.1.0-cap-CR-7557](https://github.com/argoproj/argo-rollouts/compare/v1.0.7...v1.1.0-cap-CR-7557) (2021-11-14)

### Chore

* publish notifications-install.yaml and rollout_cr_schema.json as part of release ([#1532](https://github.com/argoproj/argo-rollouts/issues/1532))
* temporarily disable ghcr.io image pushing ([#1530](https://github.com/argoproj/argo-rollouts/issues/1530))
* Fix istio vs validation reference errors ([#1454](https://github.com/argoproj/argo-rollouts/issues/1454))
* fix missing container in set image example text ([#1455](https://github.com/argoproj/argo-rollouts/issues/1455))
* Alphabetize USERS.md ([#1369](https://github.com/argoproj/argo-rollouts/issues/1369))
* Add logging step to workflow for failed e2e tests ([#1373](https://github.com/argoproj/argo-rollouts/issues/1373))
* skip e2e only if workflow started manually with debug enabled ([#1441](https://github.com/argoproj/argo-rollouts/issues/1441))
* Bump argo-ui version ([#1437](https://github.com/argoproj/argo-rollouts/issues/1437))
* Raname variables, import pkg for clarification ([#1313](https://github.com/argoproj/argo-rollouts/issues/1313))
* release workflow docker build context should use local path and not git context ([#1388](https://github.com/argoproj/argo-rollouts/issues/1388))
* github release action was using incorect docker cache ([#1387](https://github.com/argoproj/argo-rollouts/issues/1387))
* Update USERS.md by adding Databricks ([#1379](https://github.com/argoproj/argo-rollouts/issues/1379))
* add liveness and readiness probe to the install manifests ([#1324](https://github.com/argoproj/argo-rollouts/issues/1324))
* **deps:** bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 ([#1435](https://github.com/argoproj/argo-rollouts/issues/1435))
* **deps:** bump codecov/codecov-action from 1 to 2.0.3 ([#1446](https://github.com/argoproj/argo-rollouts/issues/1446))

### Ci

* add debug step to workflow ([#1374](https://github.com/argoproj/argo-rollouts/issues/1374))

### Docs

* update CHANGELOG.md for v1.1 release ([#1533](https://github.com/argoproj/argo-rollouts/issues/1533))
* clarify the service in notificatioin.md ([#1546](https://github.com/argoproj/argo-rollouts/issues/1546))
* document shell completion feature ([#1509](https://github.com/argoproj/argo-rollouts/issues/1509))
* make it clear that kustomize is required for unit tests ([#1484](https://github.com/argoproj/argo-rollouts/issues/1484))
* experiment step with traffic routing ([#1469](https://github.com/argoproj/argo-rollouts/issues/1469))
* fix missing link in toc ([#1464](https://github.com/argoproj/argo-rollouts/issues/1464))
* analysis results - handling empty array ([#1481](https://github.com/argoproj/argo-rollouts/issues/1481))
* fix typo in installation.md ([#1450](https://github.com/argoproj/argo-rollouts/issues/1450))
* Add some more explanation around TLS routes in Istio-based Rollouts README ([#1414](https://github.com/argoproj/argo-rollouts/issues/1414))
* small typo ([#1428](https://github.com/argoproj/argo-rollouts/issues/1428))
* add [@huikang](https://github.com/huikang) as a reviewer ([#1432](https://github.com/argoproj/argo-rollouts/issues/1432))
* clarify analysis. Fixes [#1234](https://github.com/argoproj/argo-rollouts/issues/1234) ([#1400](https://github.com/argoproj/argo-rollouts/issues/1400))
* clarify the setCanaryScale bahavior ([#1424](https://github.com/argoproj/argo-rollouts/issues/1424))
* Clarify quay releases. Fixes [#1408](https://github.com/argoproj/argo-rollouts/issues/1408) ([#1421](https://github.com/argoproj/argo-rollouts/issues/1421))
* add custom namespace name tips ([#1354](https://github.com/argoproj/argo-rollouts/issues/1354))
* minor grammar corrections ([#1363](https://github.com/argoproj/argo-rollouts/issues/1363))
* Add Alibaba and Ant Group to USERS.md ([#1357](https://github.com/argoproj/argo-rollouts/issues/1357))
* Link examples in documentation. Fixes [#1233](https://github.com/argoproj/argo-rollouts/issues/1233) ([#1279](https://github.com/argoproj/argo-rollouts/issues/1279))
* Preview docs locally. Fixes [#1319](https://github.com/argoproj/argo-rollouts/issues/1319) ([#1320](https://github.com/argoproj/argo-rollouts/issues/1320))
* update users list ([#1307](https://github.com/argoproj/argo-rollouts/issues/1307))
* list new metric provider integrations ([#1281](https://github.com/argoproj/argo-rollouts/issues/1281))

### Feat

* add default() evaluate helper. allow empty datadog result. Fixes [#1548](https://github.com/argoproj/argo-rollouts/issues/1548) ([#1551](https://github.com/argoproj/argo-rollouts/issues/1551))
* support dynamic scaling of stable ReplicaSet as inverse of canary weight ([#1430](https://github.com/argoproj/argo-rollouts/issues/1430))
* add support for Graphite metrics provider ([#1406](https://github.com/argoproj/argo-rollouts/issues/1406))
* create windows version for CLI ([#1517](https://github.com/argoproj/argo-rollouts/issues/1517))
* Support CloudWatch as a metric provider ([#1338](https://github.com/argoproj/argo-rollouts/issues/1338))
* provide shell completion. Closes [#619](https://github.com/argoproj/argo-rollouts/issues/619) ([#1478](https://github.com/argoproj/argo-rollouts/issues/1478))
* support management of multiple Istio VirtualService objects ([#1381](https://github.com/argoproj/argo-rollouts/issues/1381))
* Refactor dashboard code for use by Extension ([#1467](https://github.com/argoproj/argo-rollouts/issues/1467))
* kustomize rollout: add openapi to doc and examples ([#1371](https://github.com/argoproj/argo-rollouts/issues/1371))
* verify AWS TargetGroup after updating active/stable services ([#1348](https://github.com/argoproj/argo-rollouts/issues/1348))
* TrafficRouting SMI with Experiment Step in Canary ([#1351](https://github.com/argoproj/argo-rollouts/issues/1351))
* ability to abort an update when exceeding progressDeadlineSeconds ([#1397](https://github.com/argoproj/argo-rollouts/issues/1397))
* add support for Istio VirtualService spec.tls[] ([#1380](https://github.com/argoproj/argo-rollouts/issues/1380))
* configurable and more aggressive cleanup of old AnalysisRuns and Experiments ([#1342](https://github.com/argoproj/argo-rollouts/issues/1342))
* ability to auto-create Services for each template in an Experiment ([#1158](https://github.com/argoproj/argo-rollouts/issues/1158))
* introduce abortScaleDownDelaySeconds to control scale down of preview/canary upon abort ([#1160](https://github.com/argoproj/argo-rollouts/issues/1160))
* add rollout stat row to grafana dashboard ([#1343](https://github.com/argoproj/argo-rollouts/issues/1343))
* argo rollout compatibility with emissary and edge stack v2.0 ([#1330](https://github.com/argoproj/argo-rollouts/issues/1330))
* Add support for Istio multicluster ([#1274](https://github.com/argoproj/argo-rollouts/issues/1274))
* add workload-ref/generation to rollout ([#1198](https://github.com/argoproj/argo-rollouts/issues/1198))
* allow selection of namespace in rollout dashboard ([#1291](https://github.com/argoproj/argo-rollouts/issues/1291))
* support notifications on rollout events using notifications-engine ([#1175](https://github.com/argoproj/argo-rollouts/issues/1175))

### Fix

* change virtualService to pointer ([#1558](https://github.com/argoproj/argo-rollouts/issues/1558))
* Istio does not switch the traffic. Fix the VS new object creation when the tls routes nil ([#1553](https://github.com/argoproj/argo-rollouts/issues/1553))
* remove unused code of crd generation ([#1288](https://github.com/argoproj/argo-rollouts/issues/1288))
* promote nil pointer error when there are no steps ([#1510](https://github.com/argoproj/argo-rollouts/issues/1510))
* target group verification did not work with named ports ([#1485](https://github.com/argoproj/argo-rollouts/issues/1485))
* create analysisrun cmd using template generated name ([#1471](https://github.com/argoproj/argo-rollouts/issues/1471))
* replicaset count would flap when interrupting update with new pod spec ([#1479](https://github.com/argoproj/argo-rollouts/issues/1479))
* add service create to argo-rollouts role ([#1462](https://github.com/argoproj/argo-rollouts/issues/1462))
* replica count for new deployment ([#1449](https://github.com/argoproj/argo-rollouts/issues/1449))
* Nginx ingressClassName passed to canary ingress ([#1448](https://github.com/argoproj/argo-rollouts/issues/1448))
* Analysis argument validation ([#1412](https://github.com/argoproj/argo-rollouts/issues/1412))
* Remove the type from the resource name ([#1265](https://github.com/argoproj/argo-rollouts/issues/1265))
* canary scaledown event could violate maxUnavailable ([#1429](https://github.com/argoproj/argo-rollouts/issues/1429))
* analysis runs to wait for all metrics to complete ([#1407](https://github.com/argoproj/argo-rollouts/issues/1407))
* missing e2e test for istio tls route in rollout ([#1419](https://github.com/argoproj/argo-rollouts/issues/1419))
* remove unused  ServiceNotFound condition ([#1423](https://github.com/argoproj/argo-rollouts/issues/1423))
* nil pointer in create analysisrun cmd ([#1399](https://github.com/argoproj/argo-rollouts/issues/1399))
* retarget blue-green previewService before scaling up preview ReplicaSet ([#1368](https://github.com/argoproj/argo-rollouts/issues/1368))
* Promote full did not work against BlueGreen with previewReplicaCount ([#1384](https://github.com/argoproj/argo-rollouts/issues/1384))
* nil pointer dereference when reconciling paused blue-green rollout ([#1378](https://github.com/argoproj/argo-rollouts/issues/1378))
* zero-value abortScaleDownDelay was not honored with setCanaryScale ([#1375](https://github.com/argoproj/argo-rollouts/issues/1375))
* add missing notification secret to suppress error message ([#1366](https://github.com/argoproj/argo-rollouts/issues/1366))
* Abort rollout doesn't remove all canary pods for setCanaryScale ([#1352](https://github.com/argoproj/argo-rollouts/issues/1352))
* unsolicited rollout after upgrade from v0.10->v1.0 when pod was using service account ([#1367](https://github.com/argoproj/argo-rollouts/issues/1367))
* Show individual rollout in other namespace ([#1344](https://github.com/argoproj/argo-rollouts/issues/1344))
* make image using Dockerfile.dev works ([#1296](https://github.com/argoproj/argo-rollouts/issues/1296))
* abort scaledown stable RS for canary with traffic routing ([#1331](https://github.com/argoproj/argo-rollouts/issues/1331))
* lint subcommand for workload ref rollout ([#1328](https://github.com/argoproj/argo-rollouts/issues/1328))
* github action to publish docs using incorrect make target ([#1322](https://github.com/argoproj/argo-rollouts/issues/1322))
* fix mkdocs site generation. update CHANGELOG/VERSION ([#1321](https://github.com/argoproj/argo-rollouts/issues/1321))
* default replica before resolving workloadRef ([#1304](https://github.com/argoproj/argo-rollouts/issues/1304))
* missing rbac for configmaps ([#1315](https://github.com/argoproj/argo-rollouts/issues/1315))
* remove unused field in Manager struct ([#1308](https://github.com/argoproj/argo-rollouts/issues/1308))
* undo referenced object for workloadRef rollout ([#1275](https://github.com/argoproj/argo-rollouts/issues/1275))
* **metricproviders:** Check and handle invalid server URL. Fixed [#1444](https://github.com/argoproj/argo-rollouts/issues/1444) ([#1534](https://github.com/argoproj/argo-rollouts/issues/1534))
* **ui:** UI crashes on rollout view due to undefined status ([#1287](https://github.com/argoproj/argo-rollouts/issues/1287))

### Fix

* Failed analysis to degrade rollout when multiple metrics are analyzed ([#1535](https://github.com/argoproj/argo-rollouts/issues/1535))


<a name="v1.0.7"></a>
## [v1.0.7](https://github.com/argoproj/argo-rollouts/compare/v1.0.6...v1.0.7) (2021-09-23)

### Fix

* Failed analysis to degrade rollout when multiple metrics are analyzed ([#1535](https://github.com/argoproj/argo-rollouts/issues/1535))


<a name="v1.0.6"></a>
## [v1.0.6](https://github.com/argoproj/argo-rollouts/compare/v1.0.5...v1.0.6) (2021-08-25)

### Fix

* replica count for new deployment ([#1449](https://github.com/argoproj/argo-rollouts/issues/1449))
* Nginx ingressClassName passed to canary ingress ([#1448](https://github.com/argoproj/argo-rollouts/issues/1448))
* Analysis argument validation ([#1412](https://github.com/argoproj/argo-rollouts/issues/1412))


<a name="v1.0.5"></a>
## [v1.0.5](https://github.com/argoproj/argo-rollouts/compare/v1.0.4...v1.0.5) (2021-08-06)

### Chore

* release workflow docker build context should use local path and not git context ([#1388](https://github.com/argoproj/argo-rollouts/issues/1388))
* github release action was using incorect docker cache ([#1387](https://github.com/argoproj/argo-rollouts/issues/1387))

### Fix

* retarget blue-green previewService before scaling up preview ReplicaSet ([#1368](https://github.com/argoproj/argo-rollouts/issues/1368))
* analysis runs to wait for all metrics to complete ([#1407](https://github.com/argoproj/argo-rollouts/issues/1407))
* canary scaledown event could violate maxUnavailable ([#1429](https://github.com/argoproj/argo-rollouts/issues/1429))


<a name="v1.0.4"></a>
## [v1.0.4](https://github.com/argoproj/argo-rollouts/compare/v1.0.3...v1.0.4) (2021-07-30)

### Fix

* Promote full did not work against BlueGreen with previewReplicaCount ([#1384](https://github.com/argoproj/argo-rollouts/issues/1384))


<a name="v1.0.3"></a>
## [v1.0.3](https://github.com/argoproj/argo-rollouts/compare/v1.0.2...v1.0.3) (2021-07-29)

### Fix

* nil pointer dereference when reconciling paused blue-green rollout ([#1378](https://github.com/argoproj/argo-rollouts/issues/1378))
* Abort rollout doesn't remove all canary pods for setCanaryScale ([#1352](https://github.com/argoproj/argo-rollouts/issues/1352))
* unsolicited rollout after upgrade from v0.10->v1.0 when pod was using service account ([#1367](https://github.com/argoproj/argo-rollouts/issues/1367))
* default replica before resolving workloadRef ([#1304](https://github.com/argoproj/argo-rollouts/issues/1304))


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/argoproj/argo-rollouts/compare/v1.0.2-cf-init...v1.0.2) (2021-06-15)


<a name="v1.0.2-cf-init"></a>
## [v1.0.2-cf-init](https://github.com/argoproj/argo-rollouts/compare/v1.0.2-cf-6476-fix-manifests...v1.0.2-cf-init) (2021-09-26)


<a name="v1.0.2-cf-6476-fix-manifests"></a>
## [v1.0.2-cf-6476-fix-manifests](https://github.com/argoproj/argo-rollouts/compare/v1.0.1...v1.0.2-cf-6476-fix-manifests) (2021-09-26)

### Chore

* release workflow docker build context should use local path and not git context ([#1388](https://github.com/argoproj/argo-rollouts/issues/1388))
* github release action was using incorect docker cache ([#1387](https://github.com/argoproj/argo-rollouts/issues/1387))

### Docs

* always mention latest manifest. Fixes [#1262](https://github.com/argoproj/argo-rollouts/issues/1262) ([#1269](https://github.com/argoproj/argo-rollouts/issues/1269))
* Mention Helm support. Fixes [#1119](https://github.com/argoproj/argo-rollouts/issues/1119) ([#1211](https://github.com/argoproj/argo-rollouts/issues/1211))
* fix missing fields in canary spec ([#1236](https://github.com/argoproj/argo-rollouts/issues/1236))
* examples/workload-ref/workload image to dockerhub registry ([#1227](https://github.com/argoproj/argo-rollouts/issues/1227))

### Feat

* allow VirtualService HTTPRoute to be inferred if there is single route ([#1273](https://github.com/argoproj/argo-rollouts/issues/1273))
* rollouts dashboard-install manifests ([#1240](https://github.com/argoproj/argo-rollouts/issues/1240))

### Fix

* replica count for new deployment ([#1449](https://github.com/argoproj/argo-rollouts/issues/1449))
* Nginx ingressClassName passed to canary ingress ([#1448](https://github.com/argoproj/argo-rollouts/issues/1448))
* Analysis argument validation ([#1412](https://github.com/argoproj/argo-rollouts/issues/1412))
* retarget blue-green previewService before scaling up preview ReplicaSet ([#1368](https://github.com/argoproj/argo-rollouts/issues/1368))
* analysis runs to wait for all metrics to complete ([#1407](https://github.com/argoproj/argo-rollouts/issues/1407))
* canary scaledown event could violate maxUnavailable ([#1429](https://github.com/argoproj/argo-rollouts/issues/1429))
* Promote full did not work against BlueGreen with previewReplicaCount ([#1384](https://github.com/argoproj/argo-rollouts/issues/1384))
* nil pointer dereference when reconciling paused blue-green rollout ([#1378](https://github.com/argoproj/argo-rollouts/issues/1378))
* Abort rollout doesn't remove all canary pods for setCanaryScale ([#1352](https://github.com/argoproj/argo-rollouts/issues/1352))
* unsolicited rollout after upgrade from v0.10->v1.0 when pod was using service account ([#1367](https://github.com/argoproj/argo-rollouts/issues/1367))
* default replica before resolving workloadRef ([#1304](https://github.com/argoproj/argo-rollouts/issues/1304))
* rollout paused longer than progressDeadlineSeconds would briefly degrade ([#1268](https://github.com/argoproj/argo-rollouts/issues/1268))
* controller would drop fields when updating DestinationRules ([#1253](https://github.com/argoproj/argo-rollouts/issues/1253))
* Correct typo on FAQ page ([#1252](https://github.com/argoproj/argo-rollouts/issues/1252))
* the wrong panel title on the sample dashboard ([#1260](https://github.com/argoproj/argo-rollouts/issues/1260))
* avoid using root user in plugin container ([#1256](https://github.com/argoproj/argo-rollouts/issues/1256))
* analysis with multiple metrics ([#1261](https://github.com/argoproj/argo-rollouts/issues/1261))
* Mitigate the bug where items are re-added constantly to the workqueue. [#1193](https://github.com/argoproj/argo-rollouts/issues/1193) ([#1243](https://github.com/argoproj/argo-rollouts/issues/1243))
* workload rollout spec is invalid template is not empty ([#1224](https://github.com/argoproj/argo-rollouts/issues/1224))
* Fix error check in validation for AnalysisTemplates not found ([#1249](https://github.com/argoproj/argo-rollouts/issues/1249))
* make function call consistent with otherRSs definition ([#1171](https://github.com/argoproj/argo-rollouts/issues/1171))


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/argoproj/argo-rollouts/compare/v1.0.0...v1.0.1) (2021-05-25)

### Docs

* (istio.md): correct the comment in DestinationRule ([#1217](https://github.com/argoproj/argo-rollouts/issues/1217))
* minor formatting corrections. Fixes [#1113](https://github.com/argoproj/argo-rollouts/issues/1113) ([#1179](https://github.com/argoproj/argo-rollouts/issues/1179))

### Feat

* WebMetric to support string body responses ([#1212](https://github.com/argoproj/argo-rollouts/issues/1212))

### Fix

* Modify validation to check Analysis args passed through RO spec ([#1215](https://github.com/argoproj/argo-rollouts/issues/1215))
* AnalysisRun args could not be resolved from secret ([#1213](https://github.com/argoproj/argo-rollouts/issues/1213))


<a name="v1.0.0"></a>
## [v1.0.0](https://github.com/argoproj/argo-rollouts/compare/v1.0.0-rc1...v1.0.0) (2021-05-19)

### Chore

* use a base image that has a shell for the kubectl plugin image ([#1197](https://github.com/argoproj/argo-rollouts/issues/1197))
* don't run code scanning on PR push events ([#1185](https://github.com/argoproj/argo-rollouts/issues/1185))
* **deps:** bump crazy-max/ghaction-docker-meta from 2 to 3.1.0 ([#1164](https://github.com/argoproj/argo-rollouts/issues/1164))

### Docs

* clarifications in front page ([#1143](https://github.com/argoproj/argo-rollouts/issues/1143))
* fix some comments in the code and update workloadRef ([#1154](https://github.com/argoproj/argo-rollouts/issues/1154))
* Rename doc references with graphana to grafana ([#1169](https://github.com/argoproj/argo-rollouts/issues/1169))
* Add DoorDash blog post reference ([#1166](https://github.com/argoproj/argo-rollouts/issues/1166))
* update Ambassador doc link to the official website ([#1141](https://github.com/argoproj/argo-rollouts/issues/1141))
* clarifications for blue/green. Fixes [#1112](https://github.com/argoproj/argo-rollouts/issues/1112) ([#1125](https://github.com/argoproj/argo-rollouts/issues/1125))
* move workloadRef to v1.0 ([#1133](https://github.com/argoproj/argo-rollouts/issues/1133))

### Feat

* calculate rollout phase & message controller side ([#1186](https://github.com/argoproj/argo-rollouts/issues/1186))
* lint supporting rollout in multiple doc ([#1176](https://github.com/argoproj/argo-rollouts/issues/1176))
* record events for RolloutPaused and RolloutResumed ([#1178](https://github.com/argoproj/argo-rollouts/issues/1178))
* richer prometheus stats and Kubernetes events ([#1115](https://github.com/argoproj/argo-rollouts/issues/1115))
* add print version flag to rollouts-controller ([#1149](https://github.com/argoproj/argo-rollouts/issues/1149))

### Fix

* lint valid file with another empty object ([#1188](https://github.com/argoproj/argo-rollouts/issues/1188))
* Add edge case handling to traffic routing ([#1190](https://github.com/argoproj/argo-rollouts/issues/1190))
* kubectl rollout status --timeout should take a duration string ([#1194](https://github.com/argoproj/argo-rollouts/issues/1194))
* set rollout status to degraded for non-existent deploy  ([#1136](https://github.com/argoproj/argo-rollouts/issues/1136))
* deprecate skip-current-step command ([#1156](https://github.com/argoproj/argo-rollouts/issues/1156))
* build 'latest' tag from master. prefix 'v' in release builds ([#1132](https://github.com/argoproj/argo-rollouts/issues/1132))
* unhandled error patchVirtualService ([#1168](https://github.com/argoproj/argo-rollouts/issues/1168))
* handling error on f.close ([#1167](https://github.com/argoproj/argo-rollouts/issues/1167))
* Improve and refactor validation for AnalysisTemplates ([#1117](https://github.com/argoproj/argo-rollouts/issues/1117))
* use fixed size int32 ([#1155](https://github.com/argoproj/argo-rollouts/issues/1155))
* fix mismatch service selected lables to the deployment ([#1140](https://github.com/argoproj/argo-rollouts/issues/1140))
* better handle the deletion of ambassador canary mapping ([#1128](https://github.com/argoproj/argo-rollouts/issues/1128))
* **ui:** Dashboard crashes when rollout has no containers ([#1146](https://github.com/argoproj/argo-rollouts/issues/1146))
* **ui:** Dashboard Promote-Full is doing normal promotion ([#1139](https://github.com/argoproj/argo-rollouts/issues/1139))
* **ui:** UI crash after importing argo-ux due to multiple copies of React ([#1170](https://github.com/argoproj/argo-rollouts/issues/1170))

### Test

* improve e2e test to cover promote during an analysis run ([#1202](https://github.com/argoproj/argo-rollouts/issues/1202))


<a name="v1.0.0-rc1"></a>
## [v1.0.0-rc1](https://github.com/argoproj/argo-rollouts/compare/v0.10.2...v1.0.0-rc1) (2021-04-29)

### Chore

* fix typo objet to Object ([#969](https://github.com/argoproj/argo-rollouts/issues/969))
* Add node_modules to dockerignore ([#1109](https://github.com/argoproj/argo-rollouts/issues/1109))
* publish plugin image automatically. migrate to quay.io ([#1102](https://github.com/argoproj/argo-rollouts/issues/1102))
* Add make target to publish kubectl plugin image ([#1098](https://github.com/argoproj/argo-rollouts/issues/1098))
* make codegen more consistent across environments ([#1095](https://github.com/argoproj/argo-rollouts/issues/1095))
* fix broken Deploy-Docs and Docker CI tasks ([#1083](https://github.com/argoproj/argo-rollouts/issues/1083))
* manifests should be generated using kustomize 4.0.5 to satisfy CI
* Update OWNERS ([#1037](https://github.com/argoproj/argo-rollouts/issues/1037))
* direct users to ask questions on GitHub discussions ([#894](https://github.com/argoproj/argo-rollouts/issues/894))
* add better logging on rollout sync and e2e failures
* bump version. remove deprecated fields. upgrade golang. upgrade k8s deps ([#1009](https://github.com/argoproj/argo-rollouts/issues/1009))
* update k8s dependencies to v1.20. improve logging ([#994](https://github.com/argoproj/argo-rollouts/issues/994))
* add CII badge ([#993](https://github.com/argoproj/argo-rollouts/issues/993))
* Change PR template to reflect CLA->DCO ([#979](https://github.com/argoproj/argo-rollouts/issues/979))
* Separating out USERS list ([#978](https://github.com/argoproj/argo-rollouts/issues/978))
* automate release process ([#1126](https://github.com/argoproj/argo-rollouts/issues/1126))
* **deps:** bump github.com/stretchr/testify from 1.6.1 to 1.7.0 ([#965](https://github.com/argoproj/argo-rollouts/issues/965))
* **deps:** bump github.com/prometheus/common from 0.15.0 to 0.18.0 ([#1019](https://github.com/argoproj/argo-rollouts/issues/1019))
* **deps:** bump sigs.k8s.io/controller-tools from 0.4.0 to 0.4.1 ([#859](https://github.com/argoproj/argo-rollouts/issues/859))
* **deps:** upgrade k8s to v1.19.4 and prometheus dependencies ([#857](https://github.com/argoproj/argo-rollouts/issues/857))
* **deps:** bump github.com/newrelic/newrelic-client-go ([#850](https://github.com/argoproj/argo-rollouts/issues/850))

### Docs

* mention docker image for CLI ([#1120](https://github.com/argoproj/argo-rollouts/issues/1120))
* Added more diagrams in the documentation. Fixes [#323](https://github.com/argoproj/argo-rollouts/issues/323) ([#1093](https://github.com/argoproj/argo-rollouts/issues/1093))
* Add Ambassador doc links in the menu ([#1108](https://github.com/argoproj/argo-rollouts/issues/1108))
* new blog post for canaries and prometheus ([#1088](https://github.com/argoproj/argo-rollouts/issues/1088))
* Dedicated prometheus page, fixes [#1057](https://github.com/argoproj/argo-rollouts/issues/1057) ([#1065](https://github.com/argoproj/argo-rollouts/issues/1065))
* more detailed instructions for Controller metrics ([#1062](https://github.com/argoproj/argo-rollouts/issues/1062))
* Adding Keptn Argo Rollout Tutorial Video ([#1063](https://github.com/argoproj/argo-rollouts/issues/1063))
* more community content ([#1058](https://github.com/argoproj/argo-rollouts/issues/1058))
* Fixed known identifier bug in NewRelic analysis doc ([#1050](https://github.com/argoproj/argo-rollouts/issues/1050))
* clarify relationship with Argo CD and GitOps ([#1042](https://github.com/argoproj/argo-rollouts/issues/1042))
* improve ALB docs to encourage use of root-service. remove incorrect examples
* Add PayPal to users.md ([#1031](https://github.com/argoproj/argo-rollouts/issues/1031))
* fix web link in homepage for specification. Fixes [#1028](https://github.com/argoproj/argo-rollouts/issues/1028) ([#1034](https://github.com/argoproj/argo-rollouts/issues/1034))
* fix text for resuming a rollout. Fixes [#1030](https://github.com/argoproj/argo-rollouts/issues/1030) ([#1035](https://github.com/argoproj/argo-rollouts/issues/1035))
* add Automation of Everything video to docs ([#986](https://github.com/argoproj/argo-rollouts/issues/986))
* add user Devtron labs ([#959](https://github.com/argoproj/argo-rollouts/issues/959))
* update CHANGELOG.md for v0.10.1 ([#880](https://github.com/argoproj/argo-rollouts/issues/880))
* fix broken getting-started link ([#879](https://github.com/argoproj/argo-rollouts/issues/879))
* fix broken link to getting started ([#861](https://github.com/argoproj/argo-rollouts/issues/861))
* revamp documentation ([#860](https://github.com/argoproj/argo-rollouts/issues/860))
* fix typo intraffic-management/smi ([#837](https://github.com/argoproj/argo-rollouts/issues/837))

### Feat

* istio virtualservice and rollout in different namespaces ([#1114](https://github.com/argoproj/argo-rollouts/issues/1114))
* Allow Datadog API and APP keys to be consumed from env vars ([#1073](https://github.com/argoproj/argo-rollouts/issues/1073))
* add new features to kustomize transformer configuration ([#1116](https://github.com/argoproj/argo-rollouts/issues/1116))
* support reference model for workloads ([#676](https://github.com/argoproj/argo-rollouts/issues/676)) ([#1072](https://github.com/argoproj/argo-rollouts/issues/1072))
* Implement Ambassador to be used as traffic router for canary deployments ([#1025](https://github.com/argoproj/argo-rollouts/issues/1025))
* Add RolloutCompleted condition ([#1074](https://github.com/argoproj/argo-rollouts/issues/1074))
* support scaleDownDelaySeconds in canary w/ traffic routing ([#1056](https://github.com/argoproj/argo-rollouts/issues/1056))
* Add button that links to docs to header ([#1086](https://github.com/argoproj/argo-rollouts/issues/1086))
* Argo Rollouts api-server and UI ([#1015](https://github.com/argoproj/argo-rollouts/issues/1015))
* Create RolloutPaused condition ([#1054](https://github.com/argoproj/argo-rollouts/issues/1054))
* support a custom base URL for the new relic provider ([#1053](https://github.com/argoproj/argo-rollouts/issues/1053))
* Implement rollout status command. Fixes [#596](https://github.com/argoproj/argo-rollouts/issues/596) ([#1001](https://github.com/argoproj/argo-rollouts/issues/1001))
* Wait for canary RS to have ready replicas before shifting labels ([#1022](https://github.com/argoproj/argo-rollouts/issues/1022))
* Allow user to handle NaN result in Analysis ([#977](https://github.com/argoproj/argo-rollouts/issues/977))
* support canarying using Istio DestinationRule subsets ([#985](https://github.com/argoproj/argo-rollouts/issues/985))
* metric fields can be parameterized by analysis arguments ([#901](https://github.com/argoproj/argo-rollouts/issues/901))
* Add ability to restart maxUnavailable pods to BlueGreen strategy ([#937](https://github.com/argoproj/argo-rollouts/issues/937))
* add ability to verify canary weights before advancing steps ([#957](https://github.com/argoproj/argo-rollouts/issues/957))
* support ARM builds, remove unused components in Dockerfile ([#889](https://github.com/argoproj/argo-rollouts/issues/889))
* **controller:** Add support for  ephemeral metadata on BlueGreen rollouts. Fixes [#973](https://github.com/argoproj/argo-rollouts/issues/973) ([#974](https://github.com/argoproj/argo-rollouts/issues/974))

### Fix

* ui assets were not copied properly ([#1131](https://github.com/argoproj/argo-rollouts/issues/1131))
* image tag should include 'v' prefix ([#1130](https://github.com/argoproj/argo-rollouts/issues/1130))
* release-images job did not push tags properly ([#1129](https://github.com/argoproj/argo-rollouts/issues/1129))
* AnalysisRuns could be created continuously due to non-deterministic specs ([#1118](https://github.com/argoproj/argo-rollouts/issues/1118))
* rollout status always in progressing if analysis fails ([#1099](https://github.com/argoproj/argo-rollouts/issues/1099))
* make api-proto was not generating code properly ([#1101](https://github.com/argoproj/argo-rollouts/issues/1101))
* Improve validation for AnalysisTemplates referenced by RO ([#1094](https://github.com/argoproj/argo-rollouts/issues/1094))
* spec.preserveUnknownFields must be explicitly false to allow upgrades from apiextensions.k8s.io/v1beta1 ([#1069](https://github.com/argoproj/argo-rollouts/issues/1069))
* verify analysis arguments name with those in the rollout ([#1071](https://github.com/argoproj/argo-rollouts/issues/1071))
* promise rejection error on rollouts list on initial load ([#1085](https://github.com/argoproj/argo-rollouts/issues/1085))
* Wrong image for dashboard documentation ([#1084](https://github.com/argoproj/argo-rollouts/issues/1084))
* restart was restarting too many pods when available > spec.replicas ([#856](https://github.com/argoproj/argo-rollouts/issues/856))
* calculate scale down count. ([#1047](https://github.com/argoproj/argo-rollouts/issues/1047))
* add informational exposed ports to deployment ([#1066](https://github.com/argoproj/argo-rollouts/issues/1066))
* Fixes the regression of dropping resources from argo-rollouts crds. Fixes [#1043](https://github.com/argoproj/argo-rollouts/issues/1043) ([#1044](https://github.com/argoproj/argo-rollouts/issues/1044))
* controller panic with ephemeral containers with bad resource qty ([#1055](https://github.com/argoproj/argo-rollouts/issues/1055))
* analysis template arguments validate ([#1038](https://github.com/argoproj/argo-rollouts/issues/1038))
* metrics which errored did not retry at error interval
* wavefront queries would return no datapoints. surface evaluate errors
* Clear ProgressDeadlineExceeded Condition in paused BlueGreen Rollout ([#1002](https://github.com/argoproj/argo-rollouts/issues/1002))
* blue-green rollouts could pause prematurely during prePromotionAnalysis ([#1007](https://github.com/argoproj/argo-rollouts/issues/1007))
* Set Canary Strategy default maxUnavailable to 25% ([#981](https://github.com/argoproj/argo-rollouts/issues/981))
* should use copy rs before modify fields. ([#968](https://github.com/argoproj/argo-rollouts/issues/968))
* get rollout always return not found except default namespace ([#961](https://github.com/argoproj/argo-rollouts/issues/961))
* kubectl argo create panic: runtime error: invalid memory address or nil pointer dereference
* switch pod restart to use evict API to honor PDBs
* ephemeral metadata injection was dropping metadata injected by mutating webhooks ([#906](https://github.com/argoproj/argo-rollouts/issues/906))
* requiredForCompletion did not work for an experiment started by a rollout ([#907](https://github.com/argoproj/argo-rollouts/issues/907))
* Change e2e test ingress from ALB to Nginx ([#868](https://github.com/argoproj/argo-rollouts/issues/868))
* Correct Istio VirtualService immediately ([#874](https://github.com/argoproj/argo-rollouts/issues/874))
* plugin incorrectly treated v0.9 rollout as v0.10 when it had numeric observedGeneration ([#875](https://github.com/argoproj/argo-rollouts/issues/875))
* controller was not given privileges to update status subresource ([#844](https://github.com/argoproj/argo-rollouts/issues/844))
* **build:** remove extra set of parens in echo command ([#892](https://github.com/argoproj/argo-rollouts/issues/892))
* **manifests:** Add missing RoleBinding file ([#899](https://github.com/argoproj/argo-rollouts/issues/899))
* **ui:** Minor UI fixes and enhancements + more keyboard shortcuts ([#1081](https://github.com/argoproj/argo-rollouts/issues/1081))

### Test

* add e2e test to verify rollout/experiment/analysis integration ([#876](https://github.com/argoproj/argo-rollouts/issues/876))


<a name="v0.10.2"></a>
## [v0.10.2](https://github.com/argoproj/argo-rollouts/compare/v0.10.1...v0.10.2) (2020-12-17)

### Chore

* update install manifests to v0.10.2
* add better logging on rollout sync and e2e failures
* bump VERSION to v0.10.2

### Fix

* switch pod restart to use evict API to honor PDBs
* ephemeral metadata injection was dropping metadata injected by mutating webhooks ([#906](https://github.com/argoproj/argo-rollouts/issues/906))
* requiredForCompletion did not work for an experiment started by a rollout ([#907](https://github.com/argoproj/argo-rollouts/issues/907))
* **manifests:** Add missing RoleBinding file ([#899](https://github.com/argoproj/argo-rollouts/issues/899))


<a name="v0.10.1"></a>
## [v0.10.1](https://github.com/argoproj/argo-rollouts/compare/v0.10.0...v0.10.1) (2020-12-04)

### Chore

* bump rollout version to v0.10.1
* bump version to v0.10.1

### Fix

* plugin incorrectly treated v0.9 rollout as v0.10 when it had numeric observedGeneration ([#875](https://github.com/argoproj/argo-rollouts/issues/875))
* Correct Istio VirtualService immediately ([#874](https://github.com/argoproj/argo-rollouts/issues/874))
* Change e2e test ingress from ALB to Nginx
* restart was restarting too many pods when available > spec.replicas ([#856](https://github.com/argoproj/argo-rollouts/issues/856))

### Test

* add e2e test to verify rollout/experiment/analysis integration ([#876](https://github.com/argoproj/argo-rollouts/issues/876))


<a name="v0.10.0"></a>
## [v0.10.0](https://github.com/argoproj/argo-rollouts/compare/v0.9.3...v0.10.0) (2020-11-13)

### Chore

* update manifests for v0.10.0
* bump version to v0.10.0 on master ([#667](https://github.com/argoproj/argo-rollouts/issues/667))
* enable issue templates, pr templates, and no-response bot ([#665](https://github.com/argoproj/argo-rollouts/issues/665))
* add CI check to verify codegen ([#821](https://github.com/argoproj/argo-rollouts/issues/821))
* increase e2e reliability by retrying conflict patches ([#820](https://github.com/argoproj/argo-rollouts/issues/820))
* add github action to run e2e tests in k3s ([#694](https://github.com/argoproj/argo-rollouts/issues/694))
* adds dependency to generate plugin-docs when serving locally ([#711](https://github.com/argoproj/argo-rollouts/issues/711))
* migrate off bouk/monkey to go-mpatch ([#803](https://github.com/argoproj/argo-rollouts/issues/803))
* github action to publish images to github container registry ([#703](https://github.com/argoproj/argo-rollouts/issues/703))
* add spotify as a user to readme ([#769](https://github.com/argoproj/argo-rollouts/issues/769))
* Set preserveUnknownFields: false to prevent mistake ([#708](https://github.com/argoproj/argo-rollouts/issues/708))
* Spelling ([#733](https://github.com/argoproj/argo-rollouts/issues/733))
* update k8s dependencies to v1.18.2. revert preserveUnknownFields=false ([#752](https://github.com/argoproj/argo-rollouts/issues/752))
* **deps:** bump actions/setup-python from v2.1.2 to v2 ([#753](https://github.com/argoproj/argo-rollouts/issues/753))
* **deps:** bump actions to use major versions ([#755](https://github.com/argoproj/argo-rollouts/issues/755))
* **deps:** bump github.com/sirupsen/logrus from 1.6.0 to 1.7.0 ([#766](https://github.com/argoproj/argo-rollouts/issues/766))
* **deps:** bump github.com/spf13/cobra from 1.0.0 to 1.1.1 ([#788](https://github.com/argoproj/argo-rollouts/issues/788))
* **deps:** bump github.com/antonmedv/expr from 1.4.2 to 1.8.9 ([#763](https://github.com/argoproj/argo-rollouts/issues/763))
* **deps:** bump k8s.io/kubectl from 0.16.4 to 0.19.3 ([#827](https://github.com/argoproj/argo-rollouts/issues/827))
* **deps:** bump github.com/newrelic/newrelic-client-go ([#836](https://github.com/argoproj/argo-rollouts/issues/836))
* **deps:** bump actions/cache from v2.1.0 to v2.1.1 ([#652](https://github.com/argoproj/argo-rollouts/issues/652))
* **deps:** bump github.com/valyala/fasttemplate from 1.0.1 to 1.2.1 ([#648](https://github.com/argoproj/argo-rollouts/issues/648))
* **deps:** bump github.com/spaceapegames/go-wavefront ([#638](https://github.com/argoproj/argo-rollouts/issues/638))
* **deps:** bump github.com/pkg/errors from 0.8.1 to 0.9.1 ([#631](https://github.com/argoproj/argo-rollouts/issues/631))
* **deps:** bump gopkg.in/yaml.v2 from 2.2.8 to 2.3.0 ([#622](https://github.com/argoproj/argo-rollouts/issues/622))
* **deps:** bump github.com/spf13/cobra from 0.0.5 to 0.0.7 ([#630](https://github.com/argoproj/argo-rollouts/issues/630))

### Docs

* update CHANGELOG for v0.10.0 ([#842](https://github.com/argoproj/argo-rollouts/issues/842))
* add Quipper to users ([#824](https://github.com/argoproj/argo-rollouts/issues/824))
* update changelog for v0.9 ([#819](https://github.com/argoproj/argo-rollouts/issues/819))
* add video - Canary Deployments Made Easy In Kubernetes ([#815](https://github.com/argoproj/argo-rollouts/issues/815))
* fix incorrect of yaml indentation in doc ([#791](https://github.com/argoproj/argo-rollouts/issues/791))
* Being a bit more clear on the chance of downtime when migrating Deployment->Rollout ([#776](https://github.com/argoproj/argo-rollouts/issues/776))
* add VISITS Technologies to users ([#762](https://github.com/argoproj/argo-rollouts/issues/762))
* update CHANGELOG.md for v0.9.1 ([#745](https://github.com/argoproj/argo-rollouts/issues/745))
* fix grafana values to use compatible grafana dashboard ([#716](https://github.com/argoproj/argo-rollouts/issues/716))
* fix itemize list ([#706](https://github.com/argoproj/argo-rollouts/issues/706))
* Update description of paused field in specification ([#675](https://github.com/argoproj/argo-rollouts/issues/675))
* add roadmap.md ([#697](https://github.com/argoproj/argo-rollouts/issues/697))
* Add arg with default value to example for ClusterAnalysisTemplates ([#692](https://github.com/argoproj/argo-rollouts/issues/692))
* tab issues in analysis docs ([#657](https://github.com/argoproj/argo-rollouts/issues/657))
* fix analysis template docs for prometheus examples ([#668](https://github.com/argoproj/argo-rollouts/issues/668))
* Add comment to Rollout specification ([#655](https://github.com/argoproj/argo-rollouts/issues/655))
* add Ubie to list of Users ([#646](https://github.com/argoproj/argo-rollouts/issues/646))

### Feat

* Add New Relic metricprovider ([#772](https://github.com/argoproj/argo-rollouts/issues/772))
* plugin now surfaces InvalidSpec errors and failed analysisrun messages ([#729](https://github.com/argoproj/argo-rollouts/issues/729))
* web metrics preserve data types, allow insecure tls, and make jsonPath optional ([#731](https://github.com/argoproj/argo-rollouts/issues/731))
* writeback rollout updates to informer to prevent stale data ([#726](https://github.com/argoproj/argo-rollouts/issues/726))
* Add undo command in kubectl plugin. Fixes [#575](https://github.com/argoproj/argo-rollouts/issues/575) ([#812](https://github.com/argoproj/argo-rollouts/issues/812))
* add support for valueFrom in analysis arguments. ([#797](https://github.com/argoproj/argo-rollouts/issues/797))
* **analysis:** Add Datadog metric provider. Fixes [#702](https://github.com/argoproj/argo-rollouts/issues/702) ([#705](https://github.com/argoproj/argo-rollouts/issues/705))
* **controller:** Allow setting canary weight without side-effects. Fixes [#556](https://github.com/argoproj/argo-rollouts/issues/556) ([#677](https://github.com/argoproj/argo-rollouts/issues/677))
* **controller:** set canary/stable ephemeral metadata to pods during updates ([#770](https://github.com/argoproj/argo-rollouts/issues/770))
* **controller:** use CRD status subresource ([#789](https://github.com/argoproj/argo-rollouts/issues/789))
* **controller:** add full rollout promotion (skip analysis, pause, steps) ([#828](https://github.com/argoproj/argo-rollouts/issues/828))
* **controller:** restart pods up to maxUnavailable at a time ([#833](https://github.com/argoproj/argo-rollouts/issues/833))
* **metrics:** Adding rollout_info_replicas_desired metric. Fixes [#748](https://github.com/argoproj/argo-rollouts/issues/748) ([#749](https://github.com/argoproj/argo-rollouts/issues/749))
* **plugin:** Implement kubectl argo rollouts lint ([#783](https://github.com/argoproj/argo-rollouts/issues/783))

### Fix

* analysis controller could get into a hotloop with terminated run ([#724](https://github.com/argoproj/argo-rollouts/issues/724))
* fetch secrets on-demand to fix controller boot for large clusters ([#829](https://github.com/argoproj/argo-rollouts/issues/829))
* namespaced scoped controller support ([#818](https://github.com/argoproj/argo-rollouts/issues/818))
* scaleDownDelayRevisionLimit was off by one ([#816](https://github.com/argoproj/argo-rollouts/issues/816))
* background analysis refs were not verified. requeue InvalidSpec rollouts ([#814](https://github.com/argoproj/argo-rollouts/issues/814))
* rollout kustomize transform analysis ref should use templateName instead of name ([#809](https://github.com/argoproj/argo-rollouts/issues/809))
* Change BadValue to prevent continuous loop ([#643](https://github.com/argoproj/argo-rollouts/issues/643))
* add missing log message when a controller's syncHandler returns error ([#658](https://github.com/argoproj/argo-rollouts/issues/658))
* support azure auth ([#664](https://github.com/argoproj/argo-rollouts/issues/664))
* add Failed AnalysisRun phase status to analysis_run_metric_phase and analysis_run_phase metrics. ([#618](https://github.com/argoproj/argo-rollouts/issues/618))
* make controllers tolerant to spec marshalling errors ([#666](https://github.com/argoproj/argo-rollouts/issues/666))
* add missing Service kustomize name reference in trafficRouting/alb/rootService ([#699](https://github.com/argoproj/argo-rollouts/issues/699))
* bluegreen postPromotion failure did not work properly. add e2e blue-green analysis tests ([#751](https://github.com/argoproj/argo-rollouts/issues/751))
* unavailable stable RS was not scaled down to make room for canary ([#739](https://github.com/argoproj/argo-rollouts/issues/739))
* revision annotation was off-by-one for canary rollouts ([#725](https://github.com/argoproj/argo-rollouts/issues/725))
* do not create analysisruns with initial deploy ([#722](https://github.com/argoproj/argo-rollouts/issues/722))
* panic after switching to tolerant informers ([#719](https://github.com/argoproj/argo-rollouts/issues/719))
* kubectl plugin should use dynamic client ([#834](https://github.com/argoproj/argo-rollouts/issues/834))
* Remove unreachable code at rollout/controller ([#639](https://github.com/argoproj/argo-rollouts/issues/639))
* **controller:** blue-green with analysis was broken ([#780](https://github.com/argoproj/argo-rollouts/issues/780))
* **controller:** controller did not honor maxUnavailable during rollback ([#786](https://github.com/argoproj/argo-rollouts/issues/786))
* **controller:** fix unhandled panic from malformed rollout ([#801](https://github.com/argoproj/argo-rollouts/issues/801))
* **controller:** validation should not consider privileged security context ([#802](https://github.com/argoproj/argo-rollouts/issues/802))
* **controller:** calculate available replicas from active ReplicaSet ([#757](https://github.com/argoproj/argo-rollouts/issues/757))
* **plugin:** bluegreen scaleDownDelay was delaying Healthy status. Present errors in message field ([#768](https://github.com/argoproj/argo-rollouts/issues/768))

### Fix

* correct manifest install link in release.md ([#642](https://github.com/argoproj/argo-rollouts/issues/642))

### Perf

* Create IstioVirtualServiceLister ([#656](https://github.com/argoproj/argo-rollouts/issues/656))

### Refactor

* move reconciliation logic into rolloutContext ([#704](https://github.com/argoproj/argo-rollouts/issues/704))

### Test

* print rollout on failure. enhance test to retry aborted bluegreen. remove bdd test ([#843](https://github.com/argoproj/argo-rollouts/issues/843))
* add e2e test for previewReplicaCount and bluegreen-to-canary ([#767](https://github.com/argoproj/argo-rollouts/issues/767))
* add e2e testing framework ([#691](https://github.com/argoproj/argo-rollouts/issues/691))


<a name="v0.9.3"></a>
## [v0.9.3](https://github.com/argoproj/argo-rollouts/compare/v0.9.2...v0.9.3) (2020-10-16)

### Chore

* github action to publish images to github container registry ([#703](https://github.com/argoproj/argo-rollouts/issues/703))
* bump install manifests to v0.9.3
* bump version to v0.9.3

### Fix

* scaleDownDelayRevisionLimit was off by one ([#816](https://github.com/argoproj/argo-rollouts/issues/816))
* background analysis refs were not verified. requeue InvalidSpec rollouts ([#814](https://github.com/argoproj/argo-rollouts/issues/814))
* **controller:** fix unhandled panic from malformed rollout ([#801](https://github.com/argoproj/argo-rollouts/issues/801))
* **controller:** validation should not consider privileged security context ([#802](https://github.com/argoproj/argo-rollouts/issues/802))


<a name="v0.9.2"></a>
## [v0.9.2](https://github.com/argoproj/argo-rollouts/compare/v0.9.1...v0.9.2) (2020-10-16)

### Chore

* update manifests to use v0.9.2
* bump VERSION in release-v0.9 branch to v0.9.2

### Feat

* plugin now surfaces InvalidSpec errors and failed analysisrun messages ([#729](https://github.com/argoproj/argo-rollouts/issues/729))

### Fix

* bluegreen postPromotion failure did not work properly. add e2e blue-green analysis tests ([#751](https://github.com/argoproj/argo-rollouts/issues/751))
* **controller:** controller did not honor maxUnavailable during rollback ([#786](https://github.com/argoproj/argo-rollouts/issues/786))
* **controller:** blue-green with analysis was broken ([#780](https://github.com/argoproj/argo-rollouts/issues/780))
* **controller:** calculate available replicas from active ReplicaSet ([#757](https://github.com/argoproj/argo-rollouts/issues/757))
* **plugin:** bluegreen scaleDownDelay was delaying Healthy status. Present errors in message field ([#768](https://github.com/argoproj/argo-rollouts/issues/768))

### Test

* add e2e test for previewReplicaCount and bluegreen-to-canary ([#767](https://github.com/argoproj/argo-rollouts/issues/767))


<a name="v0.9.1"></a>
## [v0.9.1](https://github.com/argoproj/argo-rollouts/compare/v0.9.0...v0.9.1) (2020-09-28)

### Chore

* bump install manifests to v0.9.1
* add github action to run e2e tests in k3s ([#694](https://github.com/argoproj/argo-rollouts/issues/694))
* bump version to v0.9.1

### Feat

* web metrics preserve data types, allow insecure tls, and make jsonPath optional ([#731](https://github.com/argoproj/argo-rollouts/issues/731))
* writeback rollout updates to informer to prevent stale data ([#726](https://github.com/argoproj/argo-rollouts/issues/726))

### Fix

* unavailable stable RS was not scaled down to make room for canary ([#739](https://github.com/argoproj/argo-rollouts/issues/739))
* do not create analysisruns with initial deploy ([#722](https://github.com/argoproj/argo-rollouts/issues/722))
* panic after switching to tolerant informers ([#719](https://github.com/argoproj/argo-rollouts/issues/719))
* analysis controller could get into a hotloop with terminated run ([#724](https://github.com/argoproj/argo-rollouts/issues/724))
* make controllers tolerant to spec marshalling errors ([#666](https://github.com/argoproj/argo-rollouts/issues/666))
* add Failed AnalysisRun phase status to analysis_run_metric_phase and analysis_run_phase metrics. ([#618](https://github.com/argoproj/argo-rollouts/issues/618))
* support azure auth ([#664](https://github.com/argoproj/argo-rollouts/issues/664))
* add missing log message when a controller's syncHandler returns error ([#658](https://github.com/argoproj/argo-rollouts/issues/658))

### Perf

* Create IstioVirtualServiceLister ([#656](https://github.com/argoproj/argo-rollouts/issues/656))

### Test

* backport test fix to allow e2e testing without instanceID
* skip test for v0.10 setCanaryScale feature
* add e2e testing framework ([#691](https://github.com/argoproj/argo-rollouts/issues/691))


<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/argoproj/argo-rollouts/compare/v0.8.3...v0.9.0) (2020-08-17)

### CI

* Add codecov github action

### Chore

* enable Dependabot v2
* Argo Rollouts CRD validation ([#539](https://github.com/argoproj/argo-rollouts/issues/539))
* **deps:** bump actions/setup-python from v2.1.1 to v2.1.2
* **deps:** bump actions/setup-python from v1 to v2.1.1
* **deps:** bump actions/setup-go from v1 to v2.1.1
* **deps:** bump github.com/sirupsen/logrus from 1.4.2 to 1.6.0
* **deps:** bump golangci/golangci-lint-action from v1 to v2.2.0
* **deps:** bump actions/setup-go from v2.1.1 to v2.1.2
* **deps:** bump actions/cache from v1 to v2.1.0

### Ci

* Run Linting as a separate job

### Doc

* Add ALB getting started guide

### Docs

* Add missing alb getting started guide to nav
* Update README. Add SMI getting-started guide ([#565](https://github.com/argoproj/argo-rollouts/issues/565))
* Correct comment ([#566](https://github.com/argoproj/argo-rollouts/issues/566))
* Fix documents ([#569](https://github.com/argoproj/argo-rollouts/issues/569))
* add istio getting started guide ([#548](https://github.com/argoproj/argo-rollouts/issues/548))
* fix backoffLimit placement in example analysis job ([#546](https://github.com/argoproj/argo-rollouts/issues/546))
* rework documentation and getting started guide ([#535](https://github.com/argoproj/argo-rollouts/issues/535))
* Create documentation for Secret Referencing Feature ([#502](https://github.com/argoproj/argo-rollouts/issues/502))
* Update Version to v0.8.2 ([#496](https://github.com/argoproj/argo-rollouts/issues/496))
* SMI Proposal ([#487](https://github.com/argoproj/argo-rollouts/issues/487))
* Update install caveat to include v1.14 ([#494](https://github.com/argoproj/argo-rollouts/issues/494))
* Add FAQ about controller behavior with new rollout ([#489](https://github.com/argoproj/argo-rollouts/issues/489))
* Add missing field to alb docs ([#488](https://github.com/argoproj/argo-rollouts/issues/488))
* fix web analysis condition examples ([#486](https://github.com/argoproj/argo-rollouts/issues/486))
* Standardize duration string in examples ([#472](https://github.com/argoproj/argo-rollouts/issues/472))

### Feat

* Add support for rootService within ALB traffic routing ([#634](https://github.com/argoproj/argo-rollouts/issues/634))
* Controller Validation for objects referenced by Rollout ([#600](https://github.com/argoproj/argo-rollouts/issues/600))
* add shortened option -A for --all-namespaces ([#615](https://github.com/argoproj/argo-rollouts/issues/615))
* Controller Validation ([#549](https://github.com/argoproj/argo-rollouts/issues/549))
* SMI TrafficSplit Support for Canary ([#520](https://github.com/argoproj/argo-rollouts/issues/520))

### Fix

* Change BadValue to prevent continuous loop ([#643](https://github.com/argoproj/argo-rollouts/issues/643))
* Add cluster-analysis-template-crd to autogen manifests ([#598](https://github.com/argoproj/argo-rollouts/issues/598))
* Fix ActualWeight for canary
* Populate .spec.template with default values before Rollout Validation ([#580](https://github.com/argoproj/argo-rollouts/issues/580))
* Add SMI RBAC Privileges ([#564](https://github.com/argoproj/argo-rollouts/issues/564))
* Modify arg verification to check ValueFrom ([#500](https://github.com/argoproj/argo-rollouts/issues/500))
* Ensure ALB action with weight 0 marshalls correctly ([#493](https://github.com/argoproj/argo-rollouts/issues/493))
* Add missing clusterrole for deleting pods [#490](https://github.com/argoproj/argo-rollouts/issues/490)
* Make kubectl plugin backwards compat with canary.stableRS ([#482](https://github.com/argoproj/argo-rollouts/issues/482))

### Fix

* Update mockery dependency to fix go 1.14 issue
* remove hash selector after switching from bg to canary

### Test

* BDD Test for e2e testing ([#585](https://github.com/argoproj/argo-rollouts/issues/585))


<a name="v0.8.3"></a>
## [v0.8.3](https://github.com/argoproj/argo-rollouts/compare/v0.8.2...v0.8.3) (2020-06-03)

### Fix

* remove hash selector after switching from bg to canary

### Fix

* Modify arg verification to check ValueFrom ([#500](https://github.com/argoproj/argo-rollouts/issues/500))


<a name="v0.8.2"></a>
## [v0.8.2](https://github.com/argoproj/argo-rollouts/compare/v0.8.1...v0.8.2) (2020-05-06)

### Fix

* Ensure ALB action with weight 0 marshalls correctly ([#493](https://github.com/argoproj/argo-rollouts/issues/493))
* Add missing clusterrole for deleting pods [#490](https://github.com/argoproj/argo-rollouts/issues/490)


<a name="v0.8.1"></a>
## [v0.8.1](https://github.com/argoproj/argo-rollouts/compare/v0.8.0...v0.8.1) (2020-04-20)

### Fix

* Make kubectl plugin backwards compat with canary.stableRS ([#482](https://github.com/argoproj/argo-rollouts/issues/482))


<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/argoproj/argo-rollouts/compare/v0.7.2...v0.8.0) (2020-04-10)

### Chore

* set kubectl flags on root command ([#456](https://github.com/argoproj/argo-rollouts/issues/456))
* Fix wrong comment about the formula of calculating the replica number ([#447](https://github.com/argoproj/argo-rollouts/issues/447))
* Add StableRS to rollout status ([#441](https://github.com/argoproj/argo-rollouts/issues/441))
* Standardize controller-gen to v0.2.5 ([#431](https://github.com/argoproj/argo-rollouts/issues/431))
* Migrate from dep to go modules ([#331](https://github.com/argoproj/argo-rollouts/issues/331))
* Add autogenerated sites/ to gitignore ([#398](https://github.com/argoproj/argo-rollouts/issues/398))

### Docs

* Add FAQ to docs ([#468](https://github.com/argoproj/argo-rollouts/issues/468))
* Anti-Affinity Documentation ([#463](https://github.com/argoproj/argo-rollouts/issues/463))
* plugin command enhancements ([#454](https://github.com/argoproj/argo-rollouts/issues/454))
* fixes and improvements ([#437](https://github.com/argoproj/argo-rollouts/issues/437))
* generate kubectl plugin docs ([#422](https://github.com/argoproj/argo-rollouts/issues/422))
* Use correct podTemplateHashValue attribute for valueFrom ([#417](https://github.com/argoproj/argo-rollouts/issues/417))
* Update README.md ([#411](https://github.com/argoproj/argo-rollouts/issues/411))
* update web metrics section ([#381](https://github.com/argoproj/argo-rollouts/issues/381))
* Correct analysis docs ([#378](https://github.com/argoproj/argo-rollouts/issues/378))

### Feat

* Improve wavefront provider ([#465](https://github.com/argoproj/argo-rollouts/issues/465))
* Add ability to restart Pods ([#453](https://github.com/argoproj/argo-rollouts/issues/453))
* Improve Prometheus metrics ([#461](https://github.com/argoproj/argo-rollouts/issues/461))
* Introduce Anti-Affinity option to rollout strategies ([#445](https://github.com/argoproj/argo-rollouts/issues/445))
* Add ALB Ingress controller support ([#444](https://github.com/argoproj/argo-rollouts/issues/444))
* Add Blue Green Post Promotion Analysis ([#442](https://github.com/argoproj/argo-rollouts/issues/442))
* Nginx canary traffic management ([#426](https://github.com/argoproj/argo-rollouts/issues/426))
* Add Pre Promotion Bluegreen Analysis ([#415](https://github.com/argoproj/argo-rollouts/issues/415))
* Allow AnalysisTemplates to reference secrets ([#420](https://github.com/argoproj/argo-rollouts/issues/420))
* pause duration as string with time unit ([#423](https://github.com/argoproj/argo-rollouts/issues/423))
* Add metrics on queues and go client http calls ([#416](https://github.com/argoproj/argo-rollouts/issues/416))
* Add more command aliases in kubectl plugin ([#414](https://github.com/argoproj/argo-rollouts/issues/414))
* Allow Rollout to specify multiples templates ([#409](https://github.com/argoproj/argo-rollouts/issues/409))
* Refactor Experiment handling of pod hashes ([#385](https://github.com/argoproj/argo-rollouts/issues/385))
* Add patchMergeKey and patchStrategy struct tags and comments ([#386](https://github.com/argoproj/argo-rollouts/issues/386))
* Show scale down time for bluegreen replicasets ([#370](https://github.com/argoproj/argo-rollouts/issues/370)) ([#382](https://github.com/argoproj/argo-rollouts/issues/382))

### Fix

* Update Role/ClusterRole for Ingress access ([#439](https://github.com/argoproj/argo-rollouts/issues/439))
* rollout transformer for pod affinity. add new v0.7 name references and testing ([#399](https://github.com/argoproj/argo-rollouts/issues/399))
* Adding ca-certificates to docker image ([#393](https://github.com/argoproj/argo-rollouts/issues/393))

### Improvement

* Surface failure reasons for Rollouts/AnalysisRuns ([#434](https://github.com/argoproj/argo-rollouts/issues/434))

### Refactor

* Perform arg substitution in Analysis controller ([#407](https://github.com/argoproj/argo-rollouts/issues/407))
* Refactor BlueGreen Strategy ([#388](https://github.com/argoproj/argo-rollouts/issues/388))


<a name="v0.7.2"></a>
## [v0.7.2](https://github.com/argoproj/argo-rollouts/compare/v0.7.1...v0.7.2) (2020-02-25)

### Fix

* Update RS if RS's annotations need to be changed ([#413](https://github.com/argoproj/argo-rollouts/issues/413))


<a name="v0.7.1"></a>
## [v0.7.1](https://github.com/argoproj/argo-rollouts/compare/v0.7.0...v0.7.1) (2020-02-10)

### Feat

* Refactor Experiment handling of pod hashes ([#385](https://github.com/argoproj/argo-rollouts/issues/385))
* Add patchMergeKey and patchStrategy struct tags and comments ([#386](https://github.com/argoproj/argo-rollouts/issues/386))


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/argoproj/argo-rollouts/compare/v0.6.3...v0.7.0) (2020-01-21)

### Bluegreen

* allow preview service/replica sets to be replaced and fix sg fault in syncReplicasOnly ([#314](https://github.com/argoproj/argo-rollouts/issues/314))

### Chore

* vendor mockery utility ([#347](https://github.com/argoproj/argo-rollouts/issues/347))

### Docs

* fix analysis template args examples ([#357](https://github.com/argoproj/argo-rollouts/issues/357))
* fix analysis template examples ([#353](https://github.com/argoproj/argo-rollouts/issues/353))

### Feat

* Implement watch for Istio resources ([#354](https://github.com/argoproj/argo-rollouts/issues/354))
* Enhance web provider  ([#368](https://github.com/argoproj/argo-rollouts/issues/368))
* Web metric provider ([#318](https://github.com/argoproj/argo-rollouts/issues/318))
* Allow controller to delay analysis ([#350](https://github.com/argoproj/argo-rollouts/issues/350))
* Allow AnalysisRun to complete an experiment ([#345](https://github.com/argoproj/argo-rollouts/issues/345))
* Istio implementation ([#341](https://github.com/argoproj/argo-rollouts/issues/341))
* Support instance ids for rollout controller segregation ([#342](https://github.com/argoproj/argo-rollouts/issues/342))
* Wavefront metric provider ([#338](https://github.com/argoproj/argo-rollouts/issues/338))

### Fix

* Fix premature scaledown ([#365](https://github.com/argoproj/argo-rollouts/issues/365))
* Ensure podHash stays on stable-svc selector ([#340](https://github.com/argoproj/argo-rollouts/issues/340))
* omitted revisionHistoryLimit was not defaulting to 10 ([#330](https://github.com/argoproj/argo-rollouts/issues/330))


<a name="v0.6.3"></a>
## [v0.6.3](https://github.com/argoproj/argo-rollouts/compare/v0.6.2...v0.6.3) (2020-01-21)

### Fix

* Fix premature scaledown ([#365](https://github.com/argoproj/argo-rollouts/issues/365))
* Ensure podHash stays on stable-svc selector ([#340](https://github.com/argoproj/argo-rollouts/issues/340))


<a name="v0.6.2"></a>
## [v0.6.2](https://github.com/argoproj/argo-rollouts/compare/v0.6.1...v0.6.2) (2019-12-16)

### Fix

* omitted revisionHistoryLimit was not defaulting to 10 ([#330](https://github.com/argoproj/argo-rollouts/issues/330))


<a name="v0.6.1"></a>
## [v0.6.1](https://github.com/argoproj/argo-rollouts/compare/v0.6.0...v0.6.1) (2019-12-05)

### Bluegreen

* allow preview service/replica sets to be replaced and fix sg fault in syncReplicasOnly ([#314](https://github.com/argoproj/argo-rollouts/issues/314))


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/argoproj/argo-rollouts/compare/v0.5.0...v0.6.0) (2019-11-17)

### Docs

* specified type for kubectl patch ([#162](https://github.com/argoproj/argo-rollouts/issues/162))


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/argoproj/argo-rollouts/compare/v0.4.2...v0.5.0) (2019-09-13)


<a name="v0.4.2"></a>
## [v0.4.2](https://github.com/argoproj/argo-rollouts/compare/v0.4.1...v0.4.2) (2019-08-19)


<a name="v0.4.1"></a>
## [v0.4.1](https://github.com/argoproj/argo-rollouts/compare/v0.4.0...v0.4.1) (2019-06-26)


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/argoproj/argo-rollouts/compare/v0.3.2...v0.4.0) (2019-06-24)


<a name="v0.3.2"></a>
## [v0.3.2](https://github.com/argoproj/argo-rollouts/compare/v0.3.1...v0.3.2) (2019-06-12)


<a name="v0.3.1"></a>
## [v0.3.1](https://github.com/argoproj/argo-rollouts/compare/v0.3.0...v0.3.1) (2019-05-18)


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/argoproj/argo-rollouts/compare/v0.2.2...v0.3.0) (2019-04-30)


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/argoproj/argo-rollouts/compare/v0.2.1...v0.2.2) (2019-04-16)


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/argoproj/argo-rollouts/compare/v0.2.0...v0.2.1) (2019-03-20)


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/argoproj/argo-rollouts/compare/v0.1.0...v0.2.0) (2019-03-01)


<a name="v0.1.0"></a>
## v0.1.0 (2019-01-31)

### Controller

* feat: WebMetric to support string body responses (#1212)
* fix: Modify validation to check Analysis args passed through RO spec (#1215)
* fix: AnalysisRun args could not be resolved from secret (#1213)


# v1.0.0

## Notable Features
* New Argo Rollouts UI available in `kubectl argo rollouts dashboard`
* Ability to reference existing Deployment workloads instead of inlining a PodTemplate at spec.template
* Richer prometheus stats and Kubernetes events
* Support for Ambassador as a canary traffic router
* Support canarying using Istio DestinationRule subsets

## Upgrade Notes

### Installation Manifests

Installation manifests are now attached as GitHub Release artifacts (as opposed to raw files checked into git)
and can be installed with the release download URL. e.g.:

```
kubectl apply -f https://github.com/argoproj/argo-rollouts/releases/download/v1.0.0/install.yaml
```

### Argo CD OutOfSync status on Rollout v1.0.0 CRDs:

Argo Rollouts v1.0 now vends apiextensions.k8s.io/v1 CustomResourceDefinitions (previously apiextensions.k8s.io/v1beta1).
Kubernetes v1 CRDs no longer supports the preservation of unknown fields in objects, and rejects
attempts to set `spec.preserveUnknownFields: true` (the previous default). In order to support a
smooth upgrade from Argo Rollouts v0.10 to v1.0, `spec.preserveUnknownFields` is explicitly set to
`false` in the manifests, despite `false` being the default, and only option in v1 CRDs. However 
this causes diffing tools (such as Argo CD) to report the manifest as OutOfSync (since K8s drops the
false field).

More information:
* https://github.com/kubernetes-sigs/controller-tools/issues/476
* https://github.com/argoproj/argo-rollouts/pull/1069

To avoid the Argo CD OutOfSync conditions, you can remove `spec.preserveUnknownFields` from the manifests
entirely *after upgrading to v1.0*.

Alternatively, you can instruct Argo CD to ignore differences using ignoreDifferences in the Application spec:

```yaml
spec:
  ignoreDifferences:
  - group: apiextensions.k8s.io
    kind: CustomResourceDefinition
    jsonPointers:
    - /spec/preserveUnknownFields
```

### Deprcation of `kubectl argo rollouts promote --skip-current-step` flag

The promote flag `--skip-current-step` which skips the current running canary step has been
deprecated and will be removed in a future release. Its logic to skipping the current step has
been merged with the existing command:

```shell
kubectl argo rollouts promote ROLLOUT
```

The `promote ROLLOUT` command can now be used to handle both the case where the rollout needs to be
unpaused, as well as to skip the currently running canary step (e.g. an analysis/experirment/pause
step).

## Changes since v0.10

### Controller
* feat: support reference model for workloads (#676) (#1072)
* feat: Implement Ambassador to be used as traffic router for canary deployments (#1025)
* feat: support canarying using Istio DestinationRule subsets (#985)
* feat: istio virtualservice and rollout in different namespaces
* feat: add ability to verify canary weights before advancing steps (#957)
* feat: support scaleDownDelaySeconds in canary w/ traffic routing (#1056)
* feat: Add ability to restart maxUnavailable pods to BlueGreen strategy (#937)
* feat(controller): Add support for ephemeral metadata on BlueGreen rollouts. Fixes #973 (#974)
* feat: Allow user to handle NaN result in Analysis (#977)
* feat: Wait for canary RS to have ready replicas before shifting labels (#1022)
* feat: Create RolloutPaused condition (#1054)
* feat: Add RolloutCompleted condition (#1074)
* feat: add print version flag to rollouts-controller
* feat: calculate rollout phase & message controller side
* fix: Fixes the regression of dropping resources from argo-rollouts crds. Fixes #1043 (#1044)
* fix: Set Canary Strategy default maxUnavailable to 25% (#981)
* fix: blue-green rollouts could pause prematurely during prePromotionAnalysis (#1007)
* fix: Clear ProgressDeadlineExceeded Condition in paused BlueGreen Rollout (#1002)
* fix: analysis template arguments validate (#1038)
* fix: calculate scale down count. (#1047)
* fix: verify analysis arguments name with those in the rollout (#1071)
* fix: rollout status always in progressing if analysis fails (#1099)
* fix: Add edge case handling to traffic routing (#1190)
* fix: unhandled error patchVirtualService (#1168)
* fix: handling error on f.close (#1167)
* fix: rollouts in middle of restart should be considered Progressing

### Analysis
* feat: metric fields can be parameterized by analysis arguments (#901)
* feat: support a custom base URL for the new relic provider (#1053)
* feat: Allow Datadog API and APP keys to be consumed from env vars (#1073)
* fix: Improve validation for AnalysisTemplates referenced by RO (#1094)
* fix: wavefront queries would return no datapoints. surface evaluate errors
* fix: metrics which errored did not retry at error interval
* fix: Improve and refactor validation for AnalysisTemplates

### Plugin
* feat: Argo Rollouts api-server and UI (#1015)
* feat: Implement rollout status command. Fixes #596 (#1001)
* feat: lint supporting rollout in multiple doc
* fix: get rollout always return not found except default namespace (#961)
* fix: create command not support namespace in yaml file (#962)
* fix: kubectl argo create panic: runtime error: invalid memory address or nil pointer dereference

### Misc
* chore: publish plugin image automatically. migrate to quay.io (#1102)
* feat: support ARM builds, remove unused components in Dockerfile (#889)
* chore: update k8s dependencies to v1.20. improve logging (#994) 
* fix: add informational exposed ports to deployment (#1066)
* chore: Outsource reusable UI components to argo-ux npm package
* fix: use fixed size int32

# v0.10.2

## Changes since v0.10.1

### Controller
* fix: switch pod restart to use evict API to honor PDBs
* fix: ephemeral metadata injection was dropping metadata injected by mutating webhooks
* fix: requiredForCompletion did not work for an experiment started by a rollout
* fix: Add missing RoleBinding file to namespace installation

# v0.10.1

## Changes since v0.10.0

### Controller
* fix: Correct Istio VirtualService immediately (#874)
* fix: restart was restarting too many pods when available > spec.replicas (#856)

### Plugin
* fix: plugin incorrectly treated v0.9 rollout as v0.10 when it had numeric observedGeneration (#875)

# v0.10.0

## Notable Features
* Ability to set canary vs. stable ephemeral metadata on rollout Pods during an update
* Support new metric providers: New Relic, Datadog
* Ability to control canary scale during an update
* Ability to restart up to maxUnavailable pods at a time for a canary rollout
* Ability to self reference rollout metadata as arguments to analysis
* Ability to fully promote blue-green and canary rollouts (skipping steps, analysis, pauses)
* kubectl-argo-rollouts plugin command to lint rollout
* kubectl-argo-rollouts plugin command to undo a rollout (same as kubectl rollout undo)

## Upgrade Notes

Rollouts v0.10 has switched to using Kubernetes [CRD Status Subresources](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#status-subresource) ([PR #789](https://github.com/argoproj/argo-rollouts/pull/789)). This feature allows the rollout controller to record the numeric `metadata.generation` into `status.observedGeneration` which provides a reliable indicator of a Rollout who's spec has (or has not yet) been observed by the controller (for example if the argo-rollouts controller was down or delayed).

A consequence of this change, is that the v0.10 rollout controller should be used with the v0.10 kubectl-argo-rollouts plugin in order to perform actions such as abort, pause, promote. Similarly, Argo CD v1.8 should be with the v0.10 rollout controller when performing those same actions. Both kubectl-argo-rollouts plugin v0.10 and Argo CD v1.8 are backwards compatible with v0.9 rollouts controller.

## Changes since v0.9

### Controller

* feat: set canary/stable ephemeral metadata to pods during updates (#770)
* feat: add support for valueFrom in analysis arguments. (#797)
* feat: Adding rollout_info_replicas_desired metric. Fixes #748 (#749)
* feat: restart pods up to maxUnavailable at a time
* feat: add full rollout promotion (skip analysis, pause, steps) 
* feat: use CRD status subresource (#789)
* feat: Allow setting canary weight without side-effects. Fixes #556 (#677)
* fix: namespaced scoped controller support (#818)
* fix: fetch secrets on-demand to fix controller boot for large clusters (#829)

### Analysis
* feat: Add New Relic metricprovider (#772)
* feat: Add Datadog metric provider. Fixes #702 (#705)

### Plugin
* feat: Implement kubectl argo rollouts lint
* feat: Add undo command in kubectl plugin. Fixes #575 (#812)
* fix: kubectl plugin should use dynamic client

### Misc
* fix: rollout kustomize transform analysis ref should use templateName instead of name (#809)
* fix: add missing Service kustomize name reference in trafficRouting/alb/rootService (#699)


# v0.9.3

## Changes since v0.9.2

### Controller
* fix: scaleDownDelayRevisionLimit was off by one (#816)
* fix: background analysis refs were not verified. requeue InvalidSpec rollouts (#814)
* fix(controller): fix unhandled panic from malformed rollout (#801)
* fix(controller): validation should not consider privileged security context (#802)

# v0.9.2

## Changes since v0.9.1

### Controller
* fix(controller): controller did not honor maxUnavailable during rollback (#786)
* fix(controller): blue-green with analysis was broken (#780)
* fix(controller): blue-green fast-tracked rollbacks would still start analysis templates
* fix(controller): prePromotionAnalysis with previewReplicaCount would pause indefinitely w/o running analysis
* fix(controller): calculate available replicas from active ReplicaSet (#757)

### Plugin
* feat(plugin): indicate the stable ReplicaSet for blue-green rollouts in plugin
* feat(plugin): plugin now surfaces InvalidSpec errors and failed analysisrun messages (#729)
* fix(plugin): bluegreen scaleDownDelay was delaying Healthy status. Present errors in message field (#768)

# v0.9.1

## Changes since v0.9.0
### General
* feat: writeback rollout updates to informer to prevent stale data (#726)
* fix: unavailable stable RS was not scaled down to make room for canary (#739)
* fix: make controllers tolerant to spec marshalling errors (#666)
* perf: Create IstioVirtualServiceLister (#656)
* fix: add missing log message when a controller's syncHandler returns error (#658)
* fix: support azure auth (#664)

### Analysis
* feat: web metrics preserve data types, allow insecure tls, and make jsonPath optional (#731)
* fix: analysis controller could get into a hotloop with terminated run (#724)
* fix: do not create analysisruns with initial deploy (#722)
* fix: add Failed AnalysisRun phase status to analysis_run_metric_phase and analysis_run_phase metrics. (#618)

# v0.9.0

## Changes since v0.8
### General
* fix: Fix various panics #603
* feat: add security context to run as non-root #498
  
### Rollouts
* feat: Controller Validation #549
* feat: Controller Validation for objects referenced by Rollout #600
* feat: Add Rollout replicas metrics (#507) #581
* feat: Add support for rootService within ALB traffic routing #634

* fix: Populate .spec.template with default values before Rollout Validation #580
* fix: Add Rollout/scale to aggregate roles #637
* Fix: remove hash selector after switching from bg to canary #515
* fix: Set the currentStepIndex to max after bg to canary #558

### Traffic Routing
* feat: SMI TrafficSplit Support for Canary #520

### Kubectl Plugin
* feat: add shortened option -A for --all-namespaces #615

### Analysis
* feat: ClusterAnalysisTemplates (Cluster scoped AnalysisTemplates) #560
* feat: Uplevel AnalysisRun status to Rollout status #578

* fix: Modify arg verification to check ValueFrom #500
* fix: Fix analysis validation to include Kayenta #545

# v0.8.3

## Changes since v0.8.2
### General
* fix: Modify arg verification to check ValueFrom (#500)
* fix: remove hash selector after switching from bg to canary (#515)

# v0.8.2

## Changes since v0.8.1
### Rollouts
* fix: Ensure ALB action with weight 0 marshalls correctly (#493)
* fix: Add missing clusterrole for deleting pods (#490)

# v0.8.1

## Changes since v0.8.0
### General
* fix: Remove validation for limits and requests (#480)

### Rollouts
* fix: Duplicate StableRS to canary.StableRS (#483)

### Kubectl plugin
* fix: Make kubectl plugin backwards compat with canary.stableRS (#482)

# v0.8.0

## Breaking Changes
* The metric `rollout_created_time` is being removed.
* The `.status.canary.stableRS` is being deprecated for `.status.stableRS`. This release has the code to handle the migration, and the Rollout spec will updated to remove `.status.stableRS` in a future release.

## Contributors
Thank you to the following contributors for their work in this release!
* cronik
* dthomson25
* duboisf
* jessesuen
* khhirani
* moensch
* nghialv
  
## Changes since v0.7.2
### General
* feat: Improve Prometheus metrics (#461)
* feat: Add metrics on queues and go client http calls (#416)
* feat: Add patchMergeKey and patchStrategy struct tags and comments (#386)
* feat: Improve removing k8s 1.18 fields (#436)
* fix: Reduce log from error to warning (#394)
* chore: Download go deps explicitly in Dockerfile (#464)
* chore: Standardize controller-gen to v0.2.5 (#431)
* chore: Migrate from dep to go modules (#331)
* chore: Add auto generated sites/ to gitignore (#398)
* docs: Add remote name to 'make release-docs' (#435)
* docs: Documentation cleanup (#437)
* docs: Add Go mod download command to contributor docs (#425)
* docs: Corrected HPA doc (#396)
* docs: Remove extra comma in docs
* docs: Update README.md (#411)

### Rollouts
* feat: Introduce Anti-Affinity option to rollout strategies (#445)
* feat: Add ability to restart Pods (#453)
* feat: Add ALB Ingress controller support (#444)
* feat: Add Nginx canary traffic management (#426)
* feat: Add BlueGreen Pre Promotion Analysis (#415)
* feat: Add BlueGreen Post Promotion Analysis (#442)
* feat: Allow Rollout to specify multiples templates (#409)
* feat: Make pause duration as string with time unit (#423)
* feat: Use managed-by annotation (#448)
* refactor: Refactor BlueGreen Strategy (#388)
* fix: Update Role/ClusterRole for Ingress access (#439)
* fix: rollout transformer for pod affinity. add new v0.7 name references and testing (#399)
* chore: Add StableRS to rollout status (#441)
* chore: Fix wrong comment about the formula of calculating the replica number (#447)

### Analysis
* feat: Improve wavefront provider (#465)
* feat: Allow AnalysisTemplates to reference secrets (#420)
* improvement: Surface failure reasons for Rollouts/AnalysisRuns (#434)
* refactor: Perform arg substitution in Analysis controller (#407)
* docs: Use correct podTemplateHashValue attribute for valueFrom (#417)
* docs: Update web metrics section (#381)
* docs: Use correct magic value in Analysis docs (#378)

### Experiments
* feat: Experiments passed duration succeed with running analysis (#392)
* feat: Allow ex to use availableAt and finishedAt as args (#400)
* refactor: Refactor Experiment handling of pod hashes (#385)

### Kubectl plugin
* feat: Show scale down time for Blue Green ReplicaSets (#370) (#382)
* feat: Add more command aliases in kubectl plugin (#414)
* chore: Set kubectl flags on root command (#456)
* docs: Generate kubectl plugin docs (#422)
* docs: Plugin command enhancements (#454)

# v0.7.2

## Changes since v0.7.1
### Rollouts
* Update RS if RS's annotations need to be changed #413

# v0.7.1

## Changes since v0.7.0

### General
* Adding ca-certificates to docker image (#393)
* Add patchMergeKey and patchStrategy struct tags and comments (#386)
* Reduce log from error to warning (#394)

### Experiments
* Allow ex to use availableAt and finishedAt as args (#400)
* Experiments passed duration succeed with running analysis (#392)
* Refactor Experiment handling of pod hashes (#385)

# v0.7.0

## Important Notices
- Please upgrade to v0.6.x before upgrading to v0.7. Pre v0.6.0 has a different pausing logic, and v0.7.0 removes the depreciated PauseStartTime field. The v0.6.x versions have a migration script that is removed in v0.7.0. 
- This release introduces an alpha implementation of Rollouts leveraging Istio resources for traffic shaping. Check out [traffic management](https://argoproj.github.io/argo-rollouts/features/traffic-management/) for more info.

## Changes since v0.6.3

### General
* Support instance ids for rollout controller segregation #342
* Remove PauseStartTime #349
* Vendor mockery utility #347
* Remove loud log message #333

### Rollouts
#### General
* Add stableService field #337

#### Traffic Routing
* Initial Istio implementation #341
* Implement watch for Istio resources #354
* Add validation to istio virtual services #355

### Kubectl Plugin
* Introduce 'kubectl argo rollouts terminate' command #297

### Analysis

#### General
* Allow controller to delay analysis #350
* Create one background analysis per revision #309
* Allow AnalysisRun to complete an experiment #345

#### Providers
* Wavefront metric provider #338
* Web metric provider #318
* Refactor common logic in providers to library #368
* Allow web provider to be parameterized #368

# v0.6.3

## Changes since v0.6.2
### Bug Fixes

* Fix premature scaledown (#365)
* Add namespace restriction to job informer (#362)
* Fix honoring autoPromotionSeconds (#360)
* Ensure podHash stays on stable-svc selector (#340)

# v0.6.2

## Changes since v0.6.1
### Bug Fixes

* omitted revisionHistoryLimit was not defaulting to 10 (#330)
* Fix panic if rollout cannot create a new RS (#328)
* Enable controller to handle panics with crashing (#328)

# v0.6.1

## Changes since v0.6.0
### Bug Fixes

- Create one background analysis per revision (#309)
- Fix Infinite loop with PreviewReplicaCount set (#308)
- Fix a delete by zero in get command (#310)
- Set StableRS hash to current if replicaset does not actually exist (#320) 
- Bluegreen: allow preview service/replica sets to be replaced and fix sg fault in syncReplicasOnly (#314)

# v0.6.0

## Important Notices
- The pause functionality was reworked in the v0.6 release. Previously, the `.spec.paused` field was used by the controller to pause rollouts. However, this was an issue for users who wanted to manually pause the rollout since the controller assumed it was the only entity that set the field. In v0.6, the controller will add a pause condition to the `.status.pauseCondition` to pause a controller instead of setting `spec.paused`. The pause condition has a start time and a reason explaining why it paused. This allows users to set the `spec.paused` field manually and let the controller respect that pause. The v0.6 controller has a migration function to convert pre v0.6 rollouts to the new pause condition. The migration function will be removed in a future release.
- In pre-v0.6 versions, the BlueGreen strategy would have the preview service point at no ReplicaSets if the new ReplicaSet was receiving traffic from the active service. V0.6 changes that behavior to make the preview service always point at the latest ReplicaSet

## Changes since v0.5.0
#### General
- Update k8s library dependencies to v1.16 (##192)

#### Rollouts

###### Enhancements
- Add Rollout Context to reconciliation loop (##205)
- Refactor pausing (##211)
- Allow User pause (##216)
- Stop progress while paused (##193)
- Add pause condition migration (##229)
- Add abort functionality (##224)
- Rollout analysis plumbing (##183)
- Add AnalysisStep for Rollouts (##188)
- Add background analysis runs for rollouts (##196)
- Clean up old Background AnalysisRuns (##246)
- Clean up Experiments and AnalysisRuns (##197)
- Add initial Experiment Step (##165)
- Make specifying replicas/duration optional in the experiment step (##241)
- Terminate experiments from previous steps (##280)
- Add Analysis to RolloutExperimentStep (##238)
- Fix TimeOut check to consider experiment/analysis steps (##278)
- Pause a rollout upon inconclusive experiment (##256)
- Abort a rollout upon a failed experiment (##256)
- Add create AnalysisRun action in clusterrole (##231)

###### Bug Fixes
- Fix nil ptr for newRS (##233)
- Fix Rollout transformer config (##247)
- Always point preview service at the newRS (##217)
- Make active service required (##235)
- Reset ProgressDeadline on retry (##282)
- Ignore old running rs for RolloutCompleted status (##218)
- Remove scale down annot after scaling down (##187)
- Renames golang field names for blueGreen/canary to eliminate two API violations (##206)

#### Experiments
Check out the [Experiment Docs](https://github.com/argoproj/argo-rollouts/blob/release-v0.6/docs/features/experiments.md) for more information.

###### Enhancements
- Refactor experiments to use a context object (##208)
- Allow selectors to be overwritten when starting experiments (##249)
- Simplify experiment replicaset names (##274)
- Integrate Experiments with Analysis (##210)

#### Bug Fixes
- Fix experiment enqueue logic (##239)
- Annotate instead of label experiment names in replicasets (##262)
- Fix issue where a replicaset name collision could cause hot loop (##236)

#### Analysis
Check out the [Analysis Docs](https://github.com/argoproj/argo-rollouts/blob/release-v0.6/docs/features/analysis.md) for more information.
###### Enhancements
- AnalysisRun AnalysisTemplate Spec (##166)
- Initial analysis controller implementation (##168)
- Integrate analysis controller with provider interfaces (##171)
- Add metric knob for maxInconclusive (##181)
- Simplify provider interfaces to set error messages (##189)
- Implement ResumeAt logic (##232)
- Define explicit args in AnalysisTemplates and simplify AnalysisRun spec (##283)
- Use a duration string instead of int to represent duration (##286)
- Truncate measurements when greater than default (10) (##191)
- Add counter for consecutiveError (##191)

###### Prometheus Provider
-  Add initial provider and Prometheus implementation (##170)
- Rename prometheus.server to address to better reflect API client interface (##207)
- Treat NaN as inconclusive on Prometheus provider (##275)

###### Job Provider
- Implement job-based metric provider (##186)
- Job metric argument substitution. Simplify metric provider interface (##268)

###### Kayenta Provider
- Initialize check in for kayenta metric provider ##284


#### Kubectl Plugin
Check out the [kubectl plugin docs](https://github.com/argoproj/argo-rollouts/blob/release-v0.6/docs/features/kubectl-plugin.md) for more information.

###### Enhancements
- Implement argo rollouts kubectl plugin (##195)
- Introduce `kubectl argo rollout list rollouts` command (##204)
- Introduce `kubectl argo rollout list experiments` command (##267)
- Introduce `kubectl argo rollout set image` command (##251)
- Introduce `kubectl argo rollout get` command (##230)
- Introduce `kubectl argo rollout promote` command (##277)
- Add ability to `kubectl argo rollouts set image *=myrepo/myimage` (##290)
- Add `get/retry experiment` commands. Support experiment retries (##263)
- Show running jobs as part of analysis runs (##278)
- Surface experiment images to CLI (##274)

# v0.5.0

## Changes since v0.4.2
#### Bug Fixes
- Rollout deletionTimestamp are not honored (##109)
- status.availableReplicas should not count old stacks (##143)
- Fix Infinitely loop on controller loop (##146)

#### Enhancements
- Fast rollback in BlueGreen during scale down period ##127
- Attach independent scaleDownDelays to older ReplicaSets ##145
- Add scaleDownDelayRevisionLimit to limit the number of old ReplicaSets scaled up ##129

#### Experiment CRD
This release of Argo Rollouts introduces the experiment CRD. The experiment CRD allows users to define multiple PodSpec's to run for a specific duration of time. This will help enable the Kayenta use-case where a user will need to start two versions of their application at the same time. Otherwise, the users cannot have an apples-to-apples comparison of these two versions as one will skew as a result of running for a longer period.

# v0.4.2

## Changes since v0.4.1
#### Bug Fixes
- Honor MaxSurge and MaxUnavailable after last step (##141)
- Fix maxSurge maxUnavailable zero check (##135)
- Add .Spec.Replicas if not set in rollouts (##125)

# v0.4.1

## Changes since v0.4.0
#### Bug Fixes
- Workaround K8s inability to properly validate 'items' subfields ##114

# v0.4.0

## Important Notes Before Upgrading
- For the BlueGreen strategy, Argo Rollouts will only pause rollouts that have the field `spec.strategy.blueGreen.autoPromotionEnabled` set to false. The default value of `autoPromotionEnabled` is true and causes the rollout to immediately promote the new version once it is available. This change was implemented to make the pausing behavior of rollouts more straight-forward and you can read more about it at ##80. Argo Rollouts v0.3.2 introduces the `autoPromotionEnabled` flag without making any behavior changes, and those behavior changes are enforced starting at v0.4.0. __In order to upgrade without any issues__, the operator should first upgrade to v0.3.2 and add the `autoPromotionEnabled` flag with the appropriate value.  Afterward, they will be safe to upgrade to v0.4.

- For the Canary Strategy, the Argo Rollouts controller stores a hash of the canary steps in the rollout status to be able to detect changes in steps. If the canary steps change during a progressing canary update, the controller will change the hash and restart the steps.  If the rollout is in a completed state, the controller will only update the hash. In v0.4.0, the controller changed how the hash of the steps was calculated, and you can read more about that at this issues: ##103.  As a result, __the operator should only upgrade Argo Rollouts to v0.4.0 when all the canary rollouts have executed all steps and have completed__. Otherwise, the controller will restart the steps it has executed.

## Changes since v0.3.2
#### Enhancements
- Add Ability to specify canaryService Service object to reach only canary pods ##91
- Simplify unintuitive automatic pause behavior for blue green strategy ##80 
- Add back service informer to handle Service recreations quicker ##71 
- Use lister instead of kubernetes api call to load service ##98
- Switch to controller-gen to generate crd with complete openapi validation spec ##84

#### Bug Fixes
- Change step hashing function to derive hash from json representation ##103
- CRD validation needs to be removed for resource requests/limits ##101
- Possible to exceed revisionHistoryLimit with canary strategy ##93
- Change in pod template hash after upgrading k8s dependencies ##88
- Controller is missing patch event privileges bug ##86
- Rollouts unprotected from invalid specs ##84 
- Fix logging fields ##97

# v0.3.2

## Important BlueGreen Strategy Change
In v0.4.0, Argo Rollouts will have a breaking change where we will only pause BlueGreen rollouts if they have a new field called `autoPromotionEnabled` under the `spec.strategy.blueGreen` set to false. If the field is not listed, the default value will be true, and the rollout will immediately promote the new Rollout. This change was introduced to address https://github.com/argoproj/argo-rollouts/issues/80. 

To prepare for v0.4.0, v0.3.2 will introduce the `autoPromotionEnabled` field, but the controller will not act on the field. As a result, you can add the `autoPromotionEnabled` field without breaking your existing rollouts.

## Enhancements
- Add autoPromotionEnabled with no behavior change 
## Fixes
- Fix controller crash caused by glog attempting to write to /tmp (##94)

# v0.3.1

## Breaking Changes
Rename autoPromoteActiveServiceDelaySeconds to autoPromotionSeconds ##77

## Changes since v0.3.0

#### Enhancements
* Switch to Scratch final image ##67

#### Fixes
* Enable fast Rollback in BlueGreen ##78 
* Respect ScaleDownDelay during non-happy path ##79
* Scale down older RS on non-happy path ##76
* Fix issue where pod template hash could be computed inconsistently ##75
* Cleanup replicasets in canary deployment ##73
* Don't requeue 404 errors ##72

# v0.3.0

## New Features
- HPA Support ##37
- Prometheus Metrics ##29 ##47 
- Introduce ProgressDeadlineSeconds ##54
- Improved Scalability ##45 

## Breaking Changes

The `status.verifyingPreview` field was depreciated and move to `spec.pause`.

#### BlueGreen Specific
- Add previewReplicaCount ##64
- Add ability to auto-promote active service ##59
- Add ScaleDownDelaySeconds ##57

## Changes since v0.2.2
#### Enchantments 
- HPA Support ##37
- Prometheus Metrics ##29
- Add previewReplicaCount ##64
- Add ProgressDeadlineSeconds ##54
- Add Invalid spec checks with regards to ProgressDeadlineSeconds ##62
- Improve eventing and metrics ##61
- Improve Available Condition ##60
- Convert Kustomize V1.0 to Kustomize v2.0 ##56
- Make Metrics port customizable ##55
- Replace gometalinter with golangci ##46
- Add support for gotestsum ##52
- Remove service informer ##45 
- Replace verifying preview with Paused ##43


#### Bug fixes
- Prevent early pause before svc change in BG ##51
- Fix aggregate roles naming collision with Argo Workflows  ##44
- Use recreate strategy for controller  ##44

# v0.2.2
Add missing events permissions to the clusterrole

# v0.2.1
Changes the following clusterroles to prevent name collision with Argo Workflows
- `argo-aggregate-to-admin` to `argo-rollouts-aggregate-to-admin`
- `argo-aggregate-to-edit` to `argo-rollouts-aggregate-to-edit`
- `argo-aggregate-to-view` to `argo-rollouts-aggregate-to-view`

# v0.2.0
- Implements the initial ReplicaSet-based Canary Strategy
- Cleans up Status fields
- Implicit understanding of rollback based on steps completion and pod hash for Blue Green and Canary

# v0.1.0
* Creates a controller that manages a rollout object that mimics a deployment object
* Declaratively offers a Blue Green Strategy by creating the replicaset from the spec and managing an active and preview service to point to the new replicaset
