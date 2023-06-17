# k8s-operators

## Overview

This guide is based on the [operator-sdk tutorial](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/).

This is purely for fun and learning, and is not intended to be used in production.

## Motivation

Developers can get used to neglecting pods, lets kick that nasty habbit passive aggressively.
We can introduce a tamagotchi to the cluster, and if the pods are not healthy, consequences happen.

## Steps

### Installation

https://sdk.operatorframework.io/docs/building-operators/golang/installation/

### Create a new project

```bash
mkdir tamagotchi-operator
cd tamagotchi-operator
operator-sdk init --domain alexanderjophus.dev --repo github.com/alexanderjophus/tamagotchi-operator
```

### Add a new API

```bash
operator-sdk create api --group tamagotchi --version v1alpha1 --kind Pet --resource --controller
```

### Make your schema in Go

```go  
// PetSpec defines the desired state of Pet
type PetSpec struct{}

// PetStatus defines the observed state of Pet
type PetStatus struct {
	// +optional
	// +kubebuilder:validation:Enum=healthy;unhealthy
	Health string `json:"health,omitempty"`
	// +optional
	// +kubebuilder:validation:Enum=small;medium;large
	Size string `json:"size,omitempty"`
}
```

Add this above the `Pet` struct. To create `additionalPrinterColumns` for the CRD.

```go
//+kubebuilder:printcolumn:name="Health",type="string",JSONPath=".status.health"
```

### Generate CRD

```bash
make generate
```

### Write code to implement the reconcile loop

```go
//+kubebuilder:rbac:groups=tamagotchi.alexanderjophus.dev,resources=pets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tamagotchi.alexanderjophus.dev,resources=pets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tamagotchi.alexanderjophus.dev,resources=pets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *PetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	pet := &tamagotchiv1alpha1.Pet{}
	existingPod := &corev1.Pod{}

	log.Info("⚡️ Event received! ⚡️")

	pet.Name = req.Name
	pet.Namespace = req.Namespace

	if err := r.Get(ctx, req.NamespacedName, existingPod); err != nil {
		if errors.IsNotFound(err) {
			log.Info("Pod not found, ignoring since object must be deleted")
			if err := r.Delete(ctx, pet); err != nil {
				log.Error(err, "Failed to delete Pet")
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Pod")
		return ctrl.Result{}, err
	}

	pet.Labels = existingPod.Labels

	for _, c := range existingPod.Status.Conditions {
		if c.Type == "Ready" {
			pet.Status.Health = "healthy"
		}
	}

	if err := r.upsertPet(ctx, log, req.NamespacedName, pet, existingPod); err != nil {
		log.Error(err, "Failed to upsert Pet")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PetReconciler) upsertPet(ctx context.Context, log logr.Logger, namespace types.NamespacedName, pet *tamagotchiv1alpha1.Pet, existingPod *corev1.Pod) error {
	if err := r.Get(ctx, namespace, pet); err != nil {
		if errors.IsNotFound(err) {
			r.updatePetStatus(pet, existingPod)
			if err := r.Client.Create(ctx, pet); err != nil {
				log.Error(err, "Failed to create Pet")
				return err
			}
			log.Info("Pet created successfully")
		} else {
			log.Error(err, "Failed to get Pet")
			return err
		}
	} else {
		r.updatePetStatus(pet, existingPod)
		if err := r.Client.Status().Update(ctx, pet); err != nil {
			log.Error(err, "Failed to update Pet")
			return err
		}
		log.Info("Pet updated successfully")
	}
	return nil
}

func (r *PetReconciler) updatePetStatus(pet *tamagotchiv1alpha1.Pet, existingPod *corev1.Pod) {
	pet.Labels = existingPod.Labels

	pet.Status.Health = "unhealthy"
	for _, c := range existingPod.Status.Conditions {
		if c.Type == "Ready" {
			pet.Status.Health = "healthy"
		}
	}

	pet.Status.Size = "large"
	cpuSize := 0
	memSize := 0
	for _, c := range existingPod.Spec.Containers {
		cpuSize += int(c.Resources.Requests.Cpu().MilliValue())
		memSize += int(c.Resources.Requests.Memory().Value())
	}
	if cpuSize < cpuMilliSmall && memSize < memorySmall {
		pet.Status.Size = "small"
	} else if cpuSize < cpuMilliMedium && memSize < memoryMedium {
		pet.Status.Size = "medium"
	}
}
```

### Run the operator locally against k8s

```bash
make run
```

### Create a couple of pods

```bash
kubectl run nginx --image=nginx
kubectl apply -f medium_pod.yaml
kubectl apply -f large_pod.yaml
```

### Watch the pets spawn in automagically

Use `-w` or some tool like `k9s` to watch the pets spawn in.

```bash
kubectl get pets -w
```