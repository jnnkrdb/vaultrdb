import { Routes } from '@angular/router';
import { VaultComponent } from './pages/vault/vault.component';
import { InternaldbComponent } from './pages/internaldb/internaldb.component';

export const routes: Routes = [
  { path: 'vault', component: VaultComponent },
  { path: 'internalDB', component: InternaldbComponent },
]