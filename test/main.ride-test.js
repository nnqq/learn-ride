import util from 'util';
import { exec as e } from 'child_process';

const exec = util.promisify(e);
const wvs = 10 ** 8;

describe('wallet test suite', async () => {
  this.timeout(100000);

  beforeEach(async () => {
    await setupAccounts({
      contract: 10 * wvs,
      owner: 10 * wvs,
      user: 10 * wvs,
    });

    const script = compile(file('ride/main.ride'));
    const ssTx = setScript({ script }, accounts.contract);
    await broadcast(ssTx);
    await waitForTx(ssTx.id);
  });

  afterEach(async () => {
    await exec('npm run reset-state');
  });

  it('Can set owner', async () => {
    const setOwner = invokeScript({
      dApp: address(accounts.contract),
      call: {
        function: 'setOwner',
      },
    }, accounts.owner);
    await broadcast(setOwner);
    await waitForTx(setOwner.id);

    const getOwner = invokeScript({
      dApp: address(accounts.contract),
      call: {
        function: 'getOwner',
      },
    }, accounts.user);
    await broadcast(getOwner);
    await waitForTx(getOwner.id);
  });

  // edit
  it('Only owner can mint', async () => {
    const iTxFoo = invokeScript({
      dApp: address(accounts.wallet),
      call: {
        function: 'withdraw',
        args: [{ type: 'integer', value: 2 * wvs }],
      },

    }, accounts.foofoofoofoofoofoofoofoofoofoofoo);

    expect(broadcast(iTxFoo)).to.be.rejectedWith('Not enough balance');
  });
});
