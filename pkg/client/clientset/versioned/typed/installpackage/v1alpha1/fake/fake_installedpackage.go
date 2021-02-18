// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/installpackage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstalledPackages implements InstalledPackageInterface
type FakeInstalledPackages struct {
	Fake *FakeInstallV1alpha1
	ns   string
}

var installedpackagesResource = schema.GroupVersionResource{Group: "install.package.carvel.dev", Version: "v1alpha1", Resource: "installedpackages"}

var installedpackagesKind = schema.GroupVersionKind{Group: "install.package.carvel.dev", Version: "v1alpha1", Kind: "InstalledPackage"}

// Get takes name of the installedPackage, and returns the corresponding installedPackage object, and an error if there is any.
func (c *FakeInstalledPackages) Get(name string, options v1.GetOptions) (result *v1alpha1.InstalledPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(installedpackagesResource, c.ns, name), &v1alpha1.InstalledPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPackage), err
}

// List takes label and field selectors, and returns the list of InstalledPackages that match those selectors.
func (c *FakeInstalledPackages) List(opts v1.ListOptions) (result *v1alpha1.InstalledPackageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(installedpackagesResource, installedpackagesKind, c.ns, opts), &v1alpha1.InstalledPackageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstalledPackageList{ListMeta: obj.(*v1alpha1.InstalledPackageList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstalledPackageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested installedPackages.
func (c *FakeInstalledPackages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(installedpackagesResource, c.ns, opts))

}

// Create takes the representation of a installedPackage and creates it.  Returns the server's representation of the installedPackage, and an error, if there is any.
func (c *FakeInstalledPackages) Create(installedPackage *v1alpha1.InstalledPackage) (result *v1alpha1.InstalledPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(installedpackagesResource, c.ns, installedPackage), &v1alpha1.InstalledPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPackage), err
}

// Update takes the representation of a installedPackage and updates it. Returns the server's representation of the installedPackage, and an error, if there is any.
func (c *FakeInstalledPackages) Update(installedPackage *v1alpha1.InstalledPackage) (result *v1alpha1.InstalledPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(installedpackagesResource, c.ns, installedPackage), &v1alpha1.InstalledPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPackage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstalledPackages) UpdateStatus(installedPackage *v1alpha1.InstalledPackage) (*v1alpha1.InstalledPackage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(installedpackagesResource, "status", c.ns, installedPackage), &v1alpha1.InstalledPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPackage), err
}

// Delete takes name of the installedPackage and deletes it. Returns an error if one occurs.
func (c *FakeInstalledPackages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(installedpackagesResource, c.ns, name), &v1alpha1.InstalledPackage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstalledPackages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(installedpackagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstalledPackageList{})
	return err
}

// Patch applies the patch and returns the patched installedPackage.
func (c *FakeInstalledPackages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstalledPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(installedpackagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstalledPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPackage), err
}