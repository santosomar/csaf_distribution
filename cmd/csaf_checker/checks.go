// This file is Free Software under the MIT License
// without warranty, see README.md and LICENSES/MIT.txt for details.
//
// SPDX-License-Identifier: MIT
//
// SPDX-FileCopyrightText: 2021 German Federal Office for Information Security (BSI) <https://www.bsi.bund.de>
// Software-Engineering: 2021 Intevation GmbH <https://intevation.de>

package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type processor struct {
	opts      *options
	redirects map[string]string
}

type check interface {
	run(*processor, string) error
	report(*processor, *Domain)
}

func newProcessor(opts *options) *processor {
	return &processor{
		opts:      opts,
		redirects: map[string]string{},
	}
}

func (p *processor) clean() {
	for k := range p.redirects {
		delete(p.redirects, k)
	}
}

func (p *processor) run(checks []check, domains []string) (*Report, error) {

	var report Report

	for _, d := range domains {
		for _, ch := range checks {
			if err := ch.run(p, d); err != nil {
				return nil, err
			}
		}
		domain := &Domain{Name: d}
		for _, ch := range checks {
			ch.report(p, domain)
		}
		report.Domains = append(report.Domains, domain)
		p.clean()
	}

	return &report, nil
}

func (p *processor) checkRedirect(r *http.Request, via []*http.Request) error {

	var path strings.Builder
	for i, v := range via {
		if i > 0 {
			path.WriteString(", ")
		}
		path.WriteString(v.URL.String())
	}
	p.redirects[r.URL.String()] = path.String()

	if len(via) > 10 {
		return errors.New("Too many redirections")
	}
	return nil
}

func (p *processor) httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: p.checkRedirect,
	}

	if p.opts.Insecure {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return &client
}

type baseCheck struct {
	num         int
	description string
	messages    []string
}

type tlsCheck struct {
	baseCheck
}

type redirectsCheck struct {
	baseCheck
}

type providerMetadataCheck struct {
	baseCheck
}

type securityCheck struct {
	baseCheck
}

type wellknownMetadataCheck struct {
	baseCheck
}

type dnsPathCheck struct {
	baseCheck
}

type oneFolderPerYearCheck struct {
	baseCheck
}

type indexCheck struct {
	baseCheck
}

type changesCheck struct {
	baseCheck
}

type directoryListingsCheck struct {
	baseCheck
}

type integrityCheck struct {
	baseCheck
}

type signaturesCheck struct {
	baseCheck
}

type publicPGPKeyCheck struct {
	baseCheck
}

func (bc *baseCheck) report(_ *processor, domain *Domain) {
	req := &Requirement{
		Num:         bc.num,
		Description: bc.description,
		Messages:    bc.messages,
	}
	domain.Requirements = append(domain.Requirements, req)
}

func (tc *tlsCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (tc *tlsCheck) report(p *processor, domain *Domain) {
	tc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (rc *redirectsCheck) run(*processor, string) error {
	return nil
}

func (rc *redirectsCheck) report(p *processor, domain *Domain) {
	keys := make([]string, len(p.redirects))
	var i int
	for k := range p.redirects {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i, k := range keys {
		keys[i] = fmt.Sprintf("Redirect %s: %s", k, p.redirects[k])
	}
	rc.baseCheck.messages = keys
	rc.baseCheck.report(p, domain)
}

func (pmdc *providerMetadataCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (pmdc *providerMetadataCheck) report(p *processor, domain *Domain) {
	pmdc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (sc *securityCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (sc *securityCheck) report(p *processor, domain *Domain) {
	sc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (wmdc *wellknownMetadataCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (wmdc *wellknownMetadataCheck) report(p *processor, domain *Domain) {
	wmdc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (dpc *dnsPathCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (dpc *dnsPathCheck) report(p *processor, domain *Domain) {
	dpc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (ofpyc *oneFolderPerYearCheck) report(p *processor, domain *Domain) {
	ofpyc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (ofpyc *oneFolderPerYearCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (ic *indexCheck) report(p *processor, domain *Domain) {
	ic.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (ic *indexCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (cc *changesCheck) report(p *processor, domain *Domain) {
	cc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (cc *changesCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (dlc *directoryListingsCheck) report(p *processor, domain *Domain) {
	dlc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (dlc *directoryListingsCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (ic *integrityCheck) report(p *processor, domain *Domain) {
	ic.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (ic *integrityCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (sc *signaturesCheck) report(p *processor, domain *Domain) {
	sc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (sc *signaturesCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}

func (ppkc *publicPGPKeyCheck) report(p *processor, domain *Domain) {
	ppkc.baseCheck.report(p, domain)
	// TODO: Implement me!
}

func (ppkc *publicPGPKeyCheck) run(*processor, string) error {
	// TODO: Implement me!
	return nil
}
