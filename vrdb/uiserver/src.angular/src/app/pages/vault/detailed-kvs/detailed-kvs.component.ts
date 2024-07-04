import { Component, EventEmitter, Input, Output, model, signal } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { KeyValueSet, NewKeyValueSet, StoreDBService } from '../../../services/vaultrdb/v1/storedb/store-db.service';
import { ENTER, COMMA } from '@angular/cdk/keycodes';
import { MatAutocompleteModule, MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';
import { MatChipInputEvent, MatChipEditedEvent, MatChipsModule } from '@angular/material/chips';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { TagsInputComponent } from '../../_sub-components/tags-input/tags-input.component';

@Component({
  selector: 'app-detailed-kvs',
  standalone: true,
  imports: [
    MatButtonModule,
    MatFormFieldModule, 
    MatChipsModule,
    MatInputModule,
    MatIconModule,
    MatAutocompleteModule,
    FormsModule,

    TagsInputComponent
  ],
  templateUrl: './detailed-kvs.component.html',
  styleUrl: './detailed-kvs.component.css'
})
export class DetailedKVSComponent {

  @Input({ required: true }) kvs!: KeyValueSet;
  @Output() changed: EventEmitter<string> = new EventEmitter();

  constructor(
    private storedbsvc: StoreDBService
  ) { 
    this.kvs.tags.forEach(element => {
      this.currentSelectedTags.update(currentSelectedTags => [...currentSelectedTags, element])
    }); 
  }

  // #####################################################  
  // enabled values for editing
  isEditable: boolean = false;
  
  // #####################################################  
  saveKVS() {
    // the editable-flag must be true, to actually edit/delete 
    // the object
    if (!this.isEditable) {
      console.log('not editable element:',this.kvs) 
      return
    }

    const updated: NewKeyValueSet = {
      key: this.kvs.key,
      value: this.kvs.value,
      tags: this.currentSelectedTags(),
      description: this.kvs.description
    }

    this.storedbsvc.KVS_Update(updated)
      .subscribe(result => {
        console.log('saved element:',this.kvs.key)
        this.kvs = result
        this.changed.emit('updated')
      })
  }

  deleteKVS() {
    // the editable-flag must be true, to actually edit/delete 
    // the object
    if (!this.isEditable) {
      console.log('not deletable element:',this.kvs)
      return
    }
    
    this.storedbsvc.KVS_Delete(this.kvs.key)
      .subscribe(_ => this.changed.emit('deleted'))
  }
  
  // #####################################################  
  // testing tags
  tagInput_onFocus() {
    if (this.isEditable) {
      console.log('focus tag input')
    }    
  }

  readonly separatorKeysCodes = [ ENTER, COMMA ] as const;
  readonly possibleTags = signal(<string[]>[]);
  readonly currentSelectedTags = signal(<string[]>[]);
  readonly currentSelectedTag = model('')

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();
    // Add our tag
    if (value && !this.kvs.tags.includes(value)) {
      this.kvs.tags.push(value)
    }
  }

  remove(tag: string): void {
    const index = this.kvs.tags.indexOf(tag)
    if (index > -1) {
      this.kvs.tags.splice(index, 1)
    }
  }

  edit(tag: string, event: MatChipEditedEvent) {
    const _value = event.value.trim();
    const value = tag.trim();

    // Remove fruit if it no longer has a name
    if (!value) {
      this.remove(tag);
      return;
    }
/*
    // Edit existing fruit
    this.devTags.update(devTags => {
      const index = devTags.indexOf(tag);
      if (index >= 0) {
        devTags[index] = value;
        return [...devTags];
      }
      return devTags;
    });
*/
  }
  
  selected(event: MatAutocompleteSelectedEvent): void {
    //this.possibleTags.update(possibleTags => [...possibleTags, event.option.viewValue]);
    //this.currentFruit.set('');
    //event.option.deselect();
  }
}
