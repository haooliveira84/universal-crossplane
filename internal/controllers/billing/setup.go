package billing

import (
	"context"

	"github.com/upbound/universal-crossplane/internal/controllers/billing/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/universal-crossplane/internal/meta"
)

// SetupAWSMarketplace adds the AWS Marketplace controller that registers this
// instance with AWS Marketplace.
func SetupAWSMarketplace(mgr ctrl.Manager, l logging.Logger) error {
	name := "aws-marketplace"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSRegion())
	if err != nil {
		return errors.Wrap(err, "cannot load default AWS config")
	}
	reg := aws.NewMarketplace(mgr.GetClient(), marketplacemetering.NewFromConfig(cfg), aws.MarketplacePublicKey)

	r := NewReconciler(mgr,
		WithLogger(l.WithValues("controller", name)),
		WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		WithRegisterer(reg),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&corev1.Secret{}).
		WithEventFilter(resource.NewPredicates(resource.IsNamed(meta.SecretNameEntitlement))).
		Complete(r)
}