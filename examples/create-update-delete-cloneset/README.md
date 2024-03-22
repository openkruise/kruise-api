# Create CloneSet

This example program demonstrates the fundamental operations for managing on
[CloneSet][1] resources,  such as `Create`, `List`, `Update` and `Delete`.

You can adopt the source code from this example to write programs that manage
other types of resources through the Kubernetes API.

## Running this example

Make sure you have a Kubernetes cluster with kruise-manager installed and `kubectl` is configured:

    kubectl get nodes

Compile this example on your workstation:

```
cd create-update-delete-cloneset
go build -o ./app
```

Now, run this application on your workstation with your local kubeconfig file:

```
./app
# or specify a kubeconfig file with flag
./app -kubeconfig=$HOME/.kube/config
```

Running this command will execute the following operations on your cluster:

1. **Create CloneSet:** This will create a 2 replica CloneSet. Verify with
   `kubectl get pods` and `kubectl get cloneset`.
2. **Update CloneSet:** This will update the CloneSet resource created in
      previous step by setting the replica count to 1 and changing the container
      image to `nginx:1.13`. You are encouraged to inspect the retry loop that
      handles conflicts. Verify the new replica count and container image with
      `kubectl describe cloneset demo`.
3. **List CloneSet:** This will retrieve CloneSet in the `default`
   namespace and print their names and replica counts.
4. **Delete CloneSet:** This will delete the CloneSet object and its
   dependent pod resource. Verify with `kubectl get cloneset`.

Each step is separated by an interactive prompt. You must hit the
<kbd>Return</kbd> key to proceed to the next step. You can use these prompts as
a break to take time to run `kubectl` and inspect the result of the operations
executed.

You should see an output like the following:

```
Created cloneset "demo-cloneset".
-> Press Return key to continue.

Updating cloneset...
Updated cloneset...
-> Press Return key to continue.

Listing clonesets in namespace "default":
 * demo-cloneset (1 replicas)
-> Press Return key to continue.

Deleting cloneset...
Deleted cloneset.
```


## Cleanup

Successfully running this program will clean the created artifacts. If you
terminate the program without completing, you can clean up the created
deployment with:

    kubectl delete cloneset demo-cloneset

## Troubleshooting

[1]: https://openkruise.io/docs/user-manuals/cloneset
