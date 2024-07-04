import { ENTER, COMMA } from '@angular/cdk/keycodes';
import { Component, Input, model, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatAutocompleteModule, MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatChipEditedEvent, MatChipInputEvent, MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-tags-input',
  standalone: true,
  imports: [
    MatFormFieldModule, 
    MatAutocompleteModule,
    MatButtonModule,
    MatInputModule,
    MatChipsModule,
    MatIconModule,
    FormsModule,
  ],
  templateUrl: './tags-input.component.html',
  styleUrl: './tags-input.component.css'
})
export class TagsInputComponent {

  // #####################################################  
  // enabled values for editing
  @Input({ required: true }) isEditable!: boolean;

  // #####################################################  
  // item handlers
  tagInput_onFocus() {
    if (this.isEditable) {
      console.log('focus tag input')
    }    
  }
  // #####################################################  
  // testing tags
  readonly separatorKeysCodes = [ ENTER, COMMA ] as const;
  readonly possibleTags = signal(<string[]>[]);
  readonly currentSelectedTags = signal(<string[]>[]);
  readonly currentSelectedTag = model('')

  add(event: MatChipInputEvent): void {
  //  const value = (event.value || '').trim();
  //  // Add our tag
  //  if (value && !this.kvs.tags.includes(value)) {
  //    this.kvs.tags.push(value)
  //  }
  }

  remove(tag: string): void {
  //  const index = this.kvs.tags.indexOf(tag)
  //  if (index > -1) {
  //    this.kvs.tags.splice(index, 1)
  //  }
  }

  edit(tag: string, event: MatChipEditedEvent) {
  //  const _value = event.value.trim();
  //  const value = tag.trim();
  //
  //  // Remove fruit if it no longer has a name
  //  if (!value) {
  //    this.remove(tag);
  //    return;
  //  } 
  //  
  //  // Edit existing fruit
  //  this.devTags.update(devTags => {
  //    const index = devTags.indexOf(tag);
  //    if (index >= 0) {
  //      devTags[index] = value;
  //      return [...devTags];
  //    }
  //    return devTags;
  //  });
  }

  selected(event: MatAutocompleteSelectedEvent): void {
  //  this.possibleTags.update(possibleTags => [...possibleTags, event.option.viewValue]);
  //  this.currentFruit.set('');
  //  event.option.deselect();
  }
}
