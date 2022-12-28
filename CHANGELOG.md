# [1.1.0](https://github.com/TnLCommunity/corndogs/compare/v1.0.3...v1.1.0) (2022-12-28)


### Bug Fixes

* add CleanUpTimedOut grpc server implementation to call store method ([073b448](https://github.com/TnLCommunity/corndogs/commit/073b44837e9d5af62ec7916856d7174be6d2c86a))
* add CleanUpTimedOut to interface and return unimplemented error in postgresstore ([30914f2](https://github.com/TnLCommunity/corndogs/commit/30914f21d476dc06935b9889ef6f5709f6340731))
* add flags to timeout for address and port ([2e36e81](https://github.com/TnLCommunity/corndogs/commit/2e36e81703bcbd22d67a8531ef923c462ac7e4a2))
* add override logic to postgres GetNextTask ([629bb2d](https://github.com/TnLCommunity/corndogs/commit/629bb2d0f35d476dbac6bd91fb0199efa5c0c338))
* add Queue field to cleanUpTimedOutRequests ([6638960](https://github.com/TnLCommunity/corndogs/commit/66389609c1ebf7e6b73ea03a05e46e59499422a0))
* comment out default test for now, start on no timeout test ([8c82c5e](https://github.com/TnLCommunity/corndogs/commit/8c82c5eb1633989c853d71e602d9fa4f2aed9dea))
* complete no timeout test, update timeout tests to handle timeout getting set to 0 when cleaned up ([314c497](https://github.com/TnLCommunity/corndogs/commit/314c497d47587433e468b46313d19ec3ecb6d5d8))
* create a new testID for every test and remove rand.Seed calls ([e23c948](https://github.com/TnLCommunity/corndogs/commit/e23c9489f119d57e4d08a1426aedb47ef2ff4cf2))
* make DefaultTimeout int64 ([1bc4399](https://github.com/TnLCommunity/corndogs/commit/1bc43995606e70729e6401c817bb7299c6b769dc))
* make testQueue a variable in timeout tests because its used a lot ([8816ae8](https://github.com/TnLCommunity/corndogs/commit/8816ae80e736637f162db3e8a1bc53fb5ec34786))
* move default UpdateTask logic to grpc implementation ([23e63c0](https://github.com/TnLCommunity/corndogs/commit/23e63c0e991e5b8d1b8f226df4e052643b126040))
* move defaults logic to grpc SubmitTask implementation, move DefaulWorkingSuffix to config ([4174489](https://github.com/TnLCommunity/corndogs/commit/417448934a989ae5af588beab896451d07e0916d))
* move GetNextTask timeout logic to grpc implementation ([c4dbd05](https://github.com/TnLCommunity/corndogs/commit/c4dbd0552b0ec2d942b3ea91a8e495c57c5f2c79))
* move the rest of GetNextTask default req logic to grpc implementation ([0a3cbca](https://github.com/TnLCommunity/corndogs/commit/0a3cbcabf40aede0977dfaacb3bebcf6a8260606))
* move things around, update TimedOut expected type ([147691c](https://github.com/TnLCommunity/corndogs/commit/147691cc8deb8940358b7b6e48971927cae75c90))
* move timeout explanation comment to grpc submittask implementation ([f312afb](https://github.com/TnLCommunity/corndogs/commit/f312afb7c97760041f3d38f7862d9d9f1d3f5e6e))
* move timeout logic out of postgres implementation ([105dcc7](https://github.com/TnLCommunity/corndogs/commit/105dcc7c35ad77c009a3d28aff2ee62ea1c86a12))
* new way to run server ([4efe797](https://github.com/TnLCommunity/corndogs/commit/4efe797ad3d66c36c7bf96a0bbe63e488d79fdb4))
* order of GreaterOrEqual check in TestTimeoutNoQueue ([bc0a954](https://github.com/TnLCommunity/corndogs/commit/bc0a954e447d5941081514c926dcb59b5fe85be8))
* print queue when queue is used ([8b16f90](https://github.com/TnLCommunity/corndogs/commit/8b16f90131bfa6f57e8b72717d3ffeb579e1a0b6))
* remove commented out default timeout test ([92f991e](https://github.com/TnLCommunity/corndogs/commit/92f991ef71557e3952b12867d1507f20cee5a2ac))
* remove server/subprocesses/timeout.go since its now a g\rpc call ([9d7f181](https://github.com/TnLCommunity/corndogs/commit/9d7f181297033acdfba2e909eb7cff640b2b457e))
* remove timeout defaults in grpc implementation ([de068cd](https://github.com/TnLCommunity/corndogs/commit/de068cdeed9d865cba27fa879fb63c9166a280aa))
* set timeout 0 on timed out tasks, only time out tasks with timeout > 0 ([a9ff735](https://github.com/TnLCommunity/corndogs/commit/a9ff735d78b6fa8ba7a29d0bdc1cb472a33bc276))
* update mod and sum files for new protos version ([f26271b](https://github.com/TnLCommunity/corndogs/commit/f26271b9499dfd6b20140e1768e73045d30db30d))
* update protos ([624bc51](https://github.com/TnLCommunity/corndogs/commit/624bc51009f2c79740b0f849a8514dbe649d9466))
* update protos version ([9f87ac5](https://github.com/TnLCommunity/corndogs/commit/9f87ac5dce081a4f6ed5a0e8af1deaa12e80a67e))
* update skaffold values for new postgres version and add timeoutCron example ([86c7eb1](https://github.com/TnLCommunity/corndogs/commit/86c7eb1861a47b82c5cd97149f3947bac8ac46c5))
* update store interface to use protos for CleanUpTimedOut ([3ef784e](https://github.com/TnLCommunity/corndogs/commit/3ef784e32777ae6f5ee8a9c4cbce63009a164f5d))
* use grpc call in timeout test ([377faad](https://github.com/TnLCommunity/corndogs/commit/377faad4d82b8f47ee6bdc5e7cba83db650e2ab5))
* use response in CleanUpTimedOut grpc implementation ([f963f12](https://github.com/TnLCommunity/corndogs/commit/f963f12fa386715455b3d495a4963d1e6ffa67e0))
* use run as default command in dockerfile but allow overriding for timeout ([3f98f84](https://github.com/TnLCommunity/corndogs/commit/3f98f842fbb39ffd74a92c2729fab6ebb8f2d9b8))


### Features

* add queue flag for timeout cmd ([a22b2f3](https://github.com/TnLCommunity/corndogs/commit/a22b2f30a6383f392a2b9c6adabea19b9b24b54e))
* add test for GetNextTask OverrideTimeout not being set and being equal to zero ([faa33fa](https://github.com/TnLCommunity/corndogs/commit/faa33fa90fec3a838daf8972cd367cd099d473ab))
* add test for GetNextTask overriding state ([5bb3dfd](https://github.com/TnLCommunity/corndogs/commit/5bb3dfd6b93421686a070ff320c8832bd0a76aec))
* add test for GetNextTask overriding with a timeout ([5786b88](https://github.com/TnLCommunity/corndogs/commit/5786b8867b815122ad5b1e8ca9e9a812361569c1))
* add test for GetNextTask overriding with no timeout ([e33200e](https://github.com/TnLCommunity/corndogs/commit/e33200e7830c8845c64982d7d541be7cd02925c3))
* add test for timeing out a specific queue ([5f066a9](https://github.com/TnLCommunity/corndogs/commit/5f066a9dbfab52b05aa542184ea544c71c37707e))
* add test for timing out multiple queues ([3de771e](https://github.com/TnLCommunity/corndogs/commit/3de771eaa8445ada13645e902599901f559be03c))
* first timeout test complete! ([3688d1e](https://github.com/TnLCommunity/corndogs/commit/3688d1e592e27b62e8c6ccb7367f53848e942c19))
* implement postgres timeout functionality ([aa8960a](https://github.com/TnLCommunity/corndogs/commit/aa8960ab7c6759f75da8f76899d70c6d6bac8b2e))
* implement timeout with queue field ([7074779](https://github.com/TnLCommunity/corndogs/commit/7074779cb582a177748087174e1151530269ff1f))
* use a cli to run the corndogs server ([b76e6ac](https://github.com/TnLCommunity/corndogs/commit/b76e6acea18cbbcd4aea8bf47ad42678215dc602))

## [1.0.3](https://github.com/TnLCommunity/corndogs/compare/v1.0.2...v1.0.3) (2022-11-23)


### Bug Fixes

* add todo comments for later ([6cf5c4e](https://github.com/TnLCommunity/corndogs/commit/6cf5c4ed6d6377a88dd5a4ca8d9af255d9985184))
* use DefaultWorkingSuffix in GetNextTask update sql ([0ec4150](https://github.com/TnLCommunity/corndogs/commit/0ec415005636abdba934d6585a01ceefaf29f2e4))

## [1.0.2](https://github.com/TnLCommunity/corndogs/compare/v1.0.1...v1.0.2) (2022-03-05)


### Bug Fixes

* portforward isnt required in test-command and only 5080 is needed in wait-for-ports ([e362515](https://github.com/TnLCommunity/corndogs/commit/e3625156481e5ac5af40c7b3e39642e15d2a17dc))
* use action-kind-test action to run skaffold and stuff ([a114dc3](https://github.com/TnLCommunity/corndogs/commit/a114dc312aa24b6e24c7c6e3ba71ed3727f8cea7))

## [1.0.1](https://github.com/TnLCommunity/corndogs/compare/v1.0.0...v1.0.1) (2022-03-05)


### Bug Fixes

* add back values-skaffold-local.yaml ([210dfaa](https://github.com/TnLCommunity/corndogs/commit/210dfaa1eb7ca7f321fe9c40b970b7723a9c0e68))
* attempt to add corndogs helm chart in pull_request workflow ([b780a61](https://github.com/TnLCommunity/corndogs/commit/b780a61d5cbf0f20a57b8ae4d2c36ed8ffc7bf1a))
* call registry tnlcommunity, not corndogs ([819e102](https://github.com/TnLCommunity/corndogs/commit/819e1020a0122e6f331329407beceaefca5b1f95))
* reference tnlcommunity/corndogs for remoteChart ([076e571](https://github.com/TnLCommunity/corndogs/commit/076e57137c79cfc5426e3278c4cdc13664de10cf))
* remove update helm chart dependencies step ([f57da27](https://github.com/TnLCommunity/corndogs/commit/f57da27ec35a667798c77e89dcd1c954e9aa9c87))
* update path to values-skaffold-local.yaml in skaffold.yaml ([e02a573](https://github.com/TnLCommunity/corndogs/commit/e02a5734161e5792b7937df414b4dd3659a3de3b))
* use remoteChart in skaffold.yaml to get corndogs chart ([2b91394](https://github.com/TnLCommunity/corndogs/commit/2b913947fc6afa30f3dad0b939a98ab51eaccf67))

# 1.0.0 (2022-02-28)


### Bug Fixes

* conventional commits should be used in PR names ([958345b](https://github.com/TnLCommunity/corndogs/commit/958345b87067eaa9584231b7436d9ec5622adc28))
* Update readme to mention using conventional commits in branch names. ([fa10bb0](https://github.com/TnLCommunity/corndogs/commit/fa10bb038d248dc0cf98f3577f4e2eeb08df6e04))
