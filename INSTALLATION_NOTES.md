➜ make submodules
Submodule 'build' (https://github.com/upbound/build) registered for path 'build'
Cloning into '/Users/joan.luque/Workspace/xoanmi/provider-jsonplaceholder/build'...
Submodule path 'build': checked out 'a6e25afa0d43da62b11af96a5d29627a52f32cd9'

export provider_name=RestApiExample

➜ make provider.prepare provider=${provider_name}
rm 'apis/sample/sample.go'
rm 'apis/sample/v1alpha1/doc.go'
rm 'apis/sample/v1alpha1/groupversion_info.go'
rm 'apis/sample/v1alpha1/mytype_types.go'
rm 'apis/sample/v1alpha1/zz_generated.deepcopy.go'
rm 'apis/sample/v1alpha1/zz_generated.managed.go'
rm 'apis/sample/v1alpha1/zz_generated.managedlist.go'
rm 'internal/controller/mytype/mytype.go'
rm 'internal/controller/mytype/mytype_test.go'
Removing INSTALLATION_NOTES.md
Removing Makefile.bak
Removing PROVIDER_CHECKLIST.md.bak
Removing README.md.bak
Removing apis/template.go.bak
Removing apis/v1alpha1/groupversion_info.go.bak
Removing apis/v1alpha1/providerconfig_types.go.bak
Removing apis/v1alpha1/providerconfigusage_types.go.bak
Removing cluster/images/provider-template/Dockerfile.bak
Removing cluster/local/integration_tests.sh.bak
Removing cmd/provider/main.go.bak
Removing examples/provider/config.yaml.bak
Removing examples/sample/mytype.yaml.bak
Removing examples/storeconfig/vault.yaml.bak
Removing go.mod.bak
Removing internal/controller/config/config.go.bak
Removing internal/controller/template.go.bak
Removing package/crds/sample.template.crossplane.io_mytypes.yaml.bak
Removing package/crds/template.crossplane.io_providerconfigs.yaml.bak
Removing package/crds/template.crossplane.io_providerconfigusages.yaml.bak
Removing package/crds/template.crossplane.io_storeconfigs.yaml.bak
Removing package/crossplane.yaml.bak