import { Component, Input, inject, signal } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { KeyValueSet } from '../../../services/vaultrdb/v1/storedb/store-db.service';
import { LiveAnnouncer } from '@angular/cdk/a11y';
import { ENTER, COMMA } from '@angular/cdk/keycodes';
import { MatChipInputEvent, MatChipEditedEvent, MatChipsModule } from '@angular/material/chips';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-detailed-kvs',
  standalone: true,
  imports: [
    MatButtonModule,
    MatFormFieldModule, 
    MatChipsModule,
    MatInputModule,
    MatIconModule,
    FormsModule
  ],
  templateUrl: './detailed-kvs.component.html',
  styleUrl: './detailed-kvs.component.css'
})
export class DetailedKVSComponent {

  @Input({ required: true }) kvs!: KeyValueSet;

  // #####################################################  
  // enabled values for editing
  isEditable: boolean = false;
  
  // #####################################################  
  saveKVS() {
    if (this.isEditable) {
      console.log('saved element:',this.kvs)
    } else {
      console.log('not editable element:',this.kvs)
    }
  }

  deleteKVS() {
    if (this.isEditable) {
      console.log('deleted element:',this.kvs)
    } else {
      console.log('not deletable element:',this.kvs)
    }
  }

  
  // #####################################################  
  // testing tags
  readonly addOnBlur = true;
  readonly separatorKeysCodes = [ ENTER, COMMA ] as const;
  readonly devTags = signal<string[]>(['this', 'is', 'under', 'construction...']);
  readonly announcer = inject(LiveAnnouncer);

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();
    // Add our tag
    if (value) {
      this.devTags.update(devTags => [...devTags, value]);
    }
    // Clear the input value
    event.chipInput!.clear();
  }

  remove(tag: string): void {
    this.devTags.update(devTags => {
      const index = devTags.indexOf(tag);
      if (index < 0) {
        return devTags;
      }
      devTags.splice(index, 1);
      this.announcer.announce(`Removed ${tag}`);
      return [...devTags];
    });
  }

  edit(tag: string, event: MatChipEditedEvent) {
    const value = event.value.trim();

    // Remove fruit if it no longer has a name
    if (!value) {
      this.remove(tag);
      return;
    }

    // Edit existing fruit
    this.devTags.update(devTags => {
      const index = devTags.indexOf(tag);
      if (index >= 0) {
        devTags[index] = value;
        return [...devTags];
      }
      return devTags;
    });
  }
}
