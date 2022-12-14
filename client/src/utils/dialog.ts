import { openDialog } from 'vue3-promise-dialog';
import ConfirmDialog from '@/components/dialogs/ConfirmDialog.vue';
import AlertDialog from '@/components/dialogs/AlertDialog.vue';

export async function confirm(text: string) {
  return await openDialog(ConfirmDialog, {text});
}

export async function alert(text: string) {
  return await openDialog(AlertDialog, {text});
}
