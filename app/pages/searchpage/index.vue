<template>
    <div class="page-wrapper">
        <div class="container-xl">
          <!-- Page title -->
          <div class="page-header d-print-none">
            <div class="row align-items-center">
              <div class="col">
                <h2 class="page-title">
                  Search results
                </h2>
                <div class="text-muted mt-1">About 2,410 result (0.19 seconds)</div>
              </div>
            </div>
          </div>
        </div>
        <div class="page-body">
          <div class="container-xl">
            <div class="row">
              <div class="col-3">
                <form action="" method="get">
                  <div class="subheader mb-2">Category</div>
                  <div class="list-group list-group-transparent mb-3">
                    <a class="list-group-item list-group-item-action d-flex align-items-center active" href="#">
                      Games
                      <small class="text-muted ms-auto">24</small>
                    </a>
                    <a class="list-group-item list-group-item-action d-flex align-items-center" href="#">
                      Clothing
                      <small class="text-muted ms-auto">149</small>
                    </a>
                    <a class="list-group-item list-group-item-action d-flex align-items-center" href="#">
                      Jewelery
                      <small class="text-muted ms-auto">88</small>
                    </a>
                    <a class="list-group-item list-group-item-action d-flex align-items-center" href="#">
                      Toys
                      <small class="text-muted ms-auto">54</small>
                    </a>
                  </div>
                  <div class="subheader mb-2">Rating</div>
                  <div class="mb-3">
                    <label class="form-check mb-1">
                      <input type="radio" class="form-check-input" name="form-stars" value="5 stars" checked="">
                      <span class="form-check-label">5 stars</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="radio" class="form-check-input" name="form-stars" value="4 stars">
                      <span class="form-check-label">4 stars</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="radio" class="form-check-input" name="form-stars" value="3 stars">
                      <span class="form-check-label">3 stars</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="radio" class="form-check-input" name="form-stars" value="2 and less stars">
                      <span class="form-check-label">2 and less stars</span>
                    </label>
                  </div>
                  <div class="subheader mb-2">Tags</div>
                  <div class="mb-3">
                    <label class="form-check mb-1">
                      <input type="checkbox" class="form-check-input" name="form-tags[]" value="business" checked="">
                      <span class="form-check-label">business</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="checkbox" class="form-check-input" name="form-tags[]" value="evening">
                      <span class="form-check-label">evening</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="checkbox" class="form-check-input" name="form-tags[]" value="leisure">
                      <span class="form-check-label">leisure</span>
                    </label>
                    <label class="form-check mb-1">
                      <input type="checkbox" class="form-check-input" name="form-tags[]" value="party">
                      <span class="form-check-label">party</span>
                    </label>
                  </div>
                  <div class="subheader mb-2">Price</div>
                  <div class="row g-2 align-items-center mb-3">
                    <div class="col">
                      <div class="input-group">
                        <span class="input-group-text">
                          $
                        </span>
                        <input type="text" class="form-control" placeholder="from" value="3" autocomplete="off">
                      </div>
                    </div>
                    <div class="col-auto">—</div>
                    <div class="col">
                      <div class="input-group">
                        <span class="input-group-text">
                          $
                        </span>
                        <input type="text" class="form-control" placeholder="to" autocomplete="off">
                      </div>
                    </div>
                  </div>
                  <div class="subheader mb-2">Shipping</div>
                  <div>
                    <select name="" class="form-select">
                      <option>United Kingdom</option>
                      <option>USA</option>
                      <option>Germany</option>
                      <option>Poland</option>
                      <option>Other…</option>
                    </select>
                  </div>
                  <div class="mt-5">
                    <button class="btn btn-primary w-100">
                      Confirm changes
                    </button>
                    <a href="#" class="btn btn-link w-100">
                      Reset to defaults
                    </a>
                  </div>
                </form>
              </div>
              
            </div>
          </div>
        </div>
        <footer class="footer footer-transparent d-print-none">
          <div class="container">
            <div class="row text-center align-items-center flex-row-reverse">
              <div class="col-lg-auto ms-lg-auto">
                <ul class="list-inline list-inline-dots mb-0">
                  <li class="list-inline-item"><a href="./docs/index.html" class="link-secondary">Documentation</a></li>
                  <li class="list-inline-item"><a href="./license.html" class="link-secondary">License</a></li>
                  <li class="list-inline-item"><a href="https://github.com/tabler/tabler" target="_blank" class="link-secondary" rel="noopener">Source code</a></li>
                  <li class="list-inline-item">
                    <a href="https://github.com/sponsors/codecalm" target="_blank" class="link-secondary" rel="noopener">
                      <!-- Download SVG icon from http://tabler-icons.io/i/heart -->
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon text-pink icon-filled icon-inline" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M19.5 13.572l-7.5 7.428l-7.5 -7.428m0 0a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"></path></svg>
                      Sponsor
                    </a>
                  </li>
                </ul>
              </div>
              <div class="col-12 col-lg-auto mt-3 mt-lg-0">
                <ul class="list-inline list-inline-dots mb-0">
                  <li class="list-inline-item">
                    Copyright © 2021
                    <a href="." class="link-secondary">Tabler</a>.
                    All rights reserved.
                  </li>
                  <li class="list-inline-item">
                    <a href="./changelog.html" class="link-secondary" rel="noopener">v1.0.0-beta3</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </footer>
      </div>
</template>